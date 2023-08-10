package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IScoreController interface {
	Answer(c echo.Context) error
	AccessAnswer(c echo.Context) error
}

type scoreController struct {
	su usecase.IScoreUsecase
	eu usecase.IEvalUsecase
}

func NewScoreController(su usecase.IScoreUsecase, eu usecase.IEvalUsecase) IScoreController {
	return &scoreController{su, eu}
}

type answerRequest struct {
	Username  string               `json:"username"`
	FileId    uint                 `json:"file_id"`
	Scorelist []model.ScoreRequest `json:"scorelist"`
}

type accessAnswerRequest struct {
	FileId uint `json:"file_id"`
}

func (sc *scoreController) Answer(c echo.Context) error {
	result := map[string]bool{"result": false}
	answer := answerRequest{}
	if err := c.Bind(&answer); err != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	username := answer.Username
	for _, score := range answer.Scorelist {
		add := model.Score{
			Username: username,
			EvalId:   score.EvalId,
			Score:    score.Score,
		}
		if err := sc.su.CreateScore(add); err != nil {
			return c.JSON(http.StatusInternalServerError, result)
		}
	}
	result["result"] = true
	return c.JSON(http.StatusOK, result)
}

func (sc *scoreController) AccessAnswer(c echo.Context) error {
	result := map[string]bool{"result": false}
	fileId := accessAnswerRequest{}
	if err := c.Bind(&fileId); err != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	evalres, err := sc.eu.GetAllEvals(uint(fileId.FileId))
	answers := []model.AnswerResponse{}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	for _, eval := range evalres {
		users, err := sc.su.GetUsersByEvalId(eval.EvalId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, result)
		}
		for _, user := range users {
			scores, err := sc.su.GetScoreByUsername(user)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, result)
			}
			addanswer := model.AnswerResponse{
				Username:          user,
				ScoreResponseList: scores,
			}
			answers = append(answers, addanswer)
		}
	}
	response := model.FinalAnswerResponse{
		Result: true,
		Answer: answers,
	}
	return c.JSON(http.StatusOK, response)
}
