package controller

import (
	"backend/model"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
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

func (fc *fileController) GetFile(c echo.Context) error {
	filesres, err := fc.fu.GetAllFiles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, filesres)
}

func (fc *fileController) FileReview(c echo.Context) error {
	fileId := c.Param("file_id")
	file_id, _ := strconv.Atoi(fileId)
	fileres, err := fc.fu.GetFileReviews(uint(file_id))
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
	filename := c.FormValue("filename")
	username := c.FormValue("username")
	file_url := c.FormValue("fileurl")
	created_at := time.Now()
	file := model.File{
		Filename:  filename,
		Username:  username,
		Fileurl:   file_url,
		CreatedAt: created_at,
	}
	file, err := fc.fu.CreateFile(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	evallist := c.FormValue("evallist")
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(evallist), &data); err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	for _, eval := range data {
		add := model.Eval{
			FileId:      file.FileId,
			Evalname:    eval["evalname"].(string),
			Evalmin:     eval["evalmin"].(uint),
			Evalmax:     eval["evalmax"].(uint),
			Explanation: eval["explanation"].(string),
		}
		if err := fc.eu.CreateEval(add); err != nil {
			return c.JSON(http.StatusInternalServerError, result)
		}
	}
	result["result"] = true
	return c.JSON(http.StatusOK, result)
}

func (fc *fileController) SearchFile(c echo.Context) error {
	username := c.FormValue("username")
	filesres, err := fc.fu.GetFileByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, filesres)
}
