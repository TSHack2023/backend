package usecase

import (
	"backend/model"
	"backend/repository"
)

type IFileUsecase interface {
	GetAllFiles() ([]model.FileResponse, error)
	GetFileReviews(fileId uint) (model.FileReviewResponse, error)
	GetFileByUsername(username string) ([]model.FileResponse, error)
	CreateFile(file model.File) (model.File, error)
}

type fileUsecase struct {
	fr repository.IFileRepository
	er repository.IEvalRepository
}

func NewFileUsecase(fr repository.IFileRepository, er repository.IEvalRepository) IFileUsecase {
	return &fileUsecase{fr, er}
}

func (fu *fileUsecase) GetAllFiles() ([]model.FileResponse, error) {
	files := []model.File{}
	if err := fu.fr.GetAllFiles(&files); err != nil {
		return nil, err
	}
	resfiles := []model.FileResponse{}
	for _, file := range files {
		resfile := model.FileResponse{
			FileId:    file.FileId,
			Filename:  file.Filename,
			Username:  file.Username,
			CreatedAt: file.CreatedAt,
		}
		resfiles = append(resfiles, resfile)
	}
	return resfiles, nil
}

func (fu *fileUsecase) GetFileReviews(fileId uint) (model.FileReviewResponse, error) {
	file := model.File{}
	if err := fu.fr.GetFileReviews(&file, fileId); err != nil {
		return model.FileReviewResponse{}, err
	}
	evals := []model.Eval{}
	if err := fu.er.GetAllEvals(&evals, fileId); err != nil {
		return model.FileReviewResponse{}, err
	}
	evalresponses := []model.EvalResponse{}
	for _, eval := range evals {
		evalres := model.EvalResponse{
			EvalId:      eval.EvalId,
			Evalname:    eval.Evalname,
			Evalmin:     eval.Evalmin,
			Evalmax:     eval.Evalmax,
			Explanation: eval.Explanation,
		}
		evalresponses = append(evalresponses, evalres)
	}
	resfilereview := model.FileReviewResponse{
		Result:   true,
		Fileurl:  file.Fileurl,
		Filename: file.Filename,
		Evallist: evalresponses,
	}
	return resfilereview, nil
}

func (fu *fileUsecase) GetFileByUsername(username string) ([]model.FileResponse, error) {
	files := []model.File{}
	if err := fu.fr.GetFileByUsername(&files, username); err != nil {
		return []model.FileResponse{}, err
	}
	resfiles := []model.FileResponse{}
	for _, file := range files {
		resfile := model.FileResponse{
			FileId:    file.FileId,
			Filename:  file.Filename,
			Username:  file.Username,
			CreatedAt: file.CreatedAt,
		}
		resfiles = append(resfiles, resfile)
	}
	return resfiles, nil
}

func (fu *fileUsecase) CreateFile(file model.File) (model.File, error) {
	if err := fu.fr.CreateFile(&file); err != nil {
		return model.File{}, err
	}
	return file, nil
}
