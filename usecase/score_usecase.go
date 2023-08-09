package usecase

import (
	"backend/model"
	"backend/repository"
)

type IScoreUsecase interface {
	GetAllScores(evalId uint) ([]model.ScoreResponse, error)
	GetScoreById(scoreId uint) (model.ScoreResponse, error)
	CreateScore(score model.Score) error
	GetScoreByUsername(username string) ([]model.ScoreResponse, error)
	GetUsersByEvalId(evalId uint) ([]string, error)
}

type scoreUsecase struct {
	sr repository.IScoreRepository
	er repository.IEvalRepository
}

func NewScoreUsecase(sr repository.IScoreRepository, er repository.IEvalRepository) IScoreUsecase {
	return &scoreUsecase{sr, er}
}

func (su *scoreUsecase) GetAllScores(evalId uint) ([]model.ScoreResponse, error) {
	scores := []model.Score{}
	if err := su.sr.GetAllScores(&scores, evalId); err != nil {
		return nil, err
	}
	resScores := []model.ScoreResponse{}
	for _, v := range scores {
		eval := model.Eval{}
		if err := su.er.GetEvalById(&eval, v.EvalId); err != nil {
			return nil, err
		}
		s := model.ScoreResponse{
			ScoreId:  v.ScoreId,
			Evalname: eval.Evalname,
			Score:    v.Score,
		}
		resScores = append(resScores, s)
	}
	return resScores, nil
}

func (su *scoreUsecase) GetScoreById(scoreId uint) (model.ScoreResponse, error) {
	score := model.Score{}
	if err := su.sr.GetScoreById(&score, scoreId); err != nil {
		return model.ScoreResponse{}, err
	}
	eval := model.Eval{}
	if err := su.er.GetEvalById(&eval, score.EvalId); err != nil {
		return model.ScoreResponse{}, err
	}
	resScore := model.ScoreResponse{
		ScoreId:  score.ScoreId,
		Evalname: eval.Evalname,
		Score:    score.Score,
	}
	return resScore, nil
}

func (su *scoreUsecase) CreateScore(score model.Score) error {
	if err := su.sr.CreateScore(&score); err != nil {
		return err
	}
	return nil
}

func (su *scoreUsecase) GetScoreByUsername(username string) ([]model.ScoreResponse, error) {
	scores := []model.Score{}
	if err := su.sr.GetScoreByUsername(&scores, username); err != nil {
		return nil, err
	}
	resScores := []model.ScoreResponse{}
	for _, v := range scores {
		eval := model.Eval{}
		if err := su.er.GetEvalById(&eval, v.EvalId); err != nil {
			return nil, err
		}
		s := model.ScoreResponse{
			ScoreId:  v.ScoreId,
			Evalname: eval.Evalname,
			Score:    v.Score,
		}
		resScores = append(resScores, s)
	}
	return resScores, nil
}

func (su *scoreUsecase) GetUsersByEvalId(evalId uint) ([]string, error) {
	users := []string{}
	if err := su.sr.GetUsersByEvalId(&users, evalId); err != nil {
		return nil, err
	}
	return users, nil
}
