package usecase

import (
	"backend/model"
	"backend/repository"
)

type IEvalUsecase interface {
	GetAllEvals(fileId uint) ([]model.EvalResponse, error)
	GetEvalById(evalId uint) (model.EvalResponse, error)
	CreateEval(eval model.Eval) error
}

type evalUsecase struct {
	er repository.IEvalRepository
}

func NewEvalUsecase(er repository.IEvalRepository) IEvalUsecase {
	return &evalUsecase{er}
}

func (eu *evalUsecase) GetAllEvals(fileId uint) ([]model.EvalResponse, error) {
	evals := []model.Eval{}
	if err := eu.er.GetAllEvals(&evals, fileId); err != nil {
		return nil, err
	}
	resEvals := []model.EvalResponse{}
	for _, v := range evals {
		e := model.EvalResponse{
			EvalId:      v.EvalId,
			Evalname:    v.Evalname,
			Evalmin:     v.Evalmin,
			Evalmax:     v.Evalmax,
			Explanation: v.Explanation,
		}
		resEvals = append(resEvals, e)
	}
	return resEvals, nil
}

func (eu *evalUsecase) GetEvalById(evalId uint) (model.EvalResponse, error) {
	eval := model.Eval{}
	if err := eu.er.GetEvalById(&eval, evalId); err != nil {
		return model.EvalResponse{}, err
	}
	resEval := model.EvalResponse{
		EvalId:      eval.EvalId,
		Evalname:    eval.Evalname,
		Evalmin:     eval.Evalmin,
		Evalmax:     eval.Evalmax,
		Explanation: eval.Explanation,
	}
	return resEval, nil
}

func (eu *evalUsecase) CreateEval(eval model.Eval) error {
	if err := eu.er.CreateEval(&eval); err != nil {
		return err
	}
	return nil
}

/*
func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
*/
