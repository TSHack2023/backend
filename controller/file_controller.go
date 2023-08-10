package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IFileController interface {
	GetFile(c echo.Context) error
	FileReview(c echo.Context) error
	FileUpload(c echo.Context) error
	SearchFile(c echo.Context) error
}

type fileController struct {
	fu usecase.IFileUsecase
	eu usecase.IEvalUsecase
}

func NewFileController(fu usecase.IFileUsecase, eu usecase.IEvalUsecase) IFileController {
	return &fileController{fu, eu}
}

type fileUploadRequest struct {
	Username string              `json:"username"`
	Filename string              `json:"filename"`
	Fileurl  string              `json:"fileurl"`
	Evallist []model.EvalRequest `json:"evallist"`
}

type searchFileRequest struct {
	Username string `json:"username"`
}

type fileReviewRequest struct {
	FileId uint `json:"file_id"`
}

func (fc *fileController) GetFile(c echo.Context) error {
	filesres, err := fc.fu.GetAllFiles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, filesres)
}

func (fc *fileController) FileReview(c echo.Context) error {
	fileId := fileReviewRequest{}
	if err := c.Bind(&fileId); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]bool{"result": false})
	}
	fileres, err := fc.fu.GetFileReviews(uint(fileId.FileId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]bool{"result": false})
	}
	response := struct {
		FileResponse model.FileReviewResponse
		Result       bool `json:"result"`
	}{
		FileResponse: fileres,
		Result:       true,
	}
	return c.JSON(http.StatusOK, response)
}

func (fc *fileController) FileUpload(c echo.Context) error {
	result := map[string]bool{"result": false}
	fileuploadrequest := fileUploadRequest{}
	if err := c.Bind(&fileuploadrequest); err != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	created_at := time.Now()
	file := model.File{
		Filename:  fileuploadrequest.Filename,
		Username:  fileuploadrequest.Username,
		Fileurl:   fileuploadrequest.Fileurl,
		CreatedAt: created_at,
	}
	file, err := fc.fu.CreateFile(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	evallist := fileuploadrequest.Evallist
	for _, eval := range evallist {
		add := model.Eval{
			FileId:      file.FileId,
			Evalname:    eval.Evalname,
			Evalmin:     eval.Evalmin,
			Evalmax:     eval.Evalmax,
			Explanation: eval.Explanation,
		}
		err := fc.eu.CreateEval(add)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, result)
		}
	}
	result["result"] = true
	return c.JSON(http.StatusOK, result)
}

func (fc *fileController) SearchFile(c echo.Context) error {
	username := searchFileRequest{}
	if err := c.Bind(&username); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	filesres, err := fc.fu.GetFileByUsername(username.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, filesres)
}
