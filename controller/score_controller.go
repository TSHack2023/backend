package controller

import (
	"backend/model"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"

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

func NewScoreController() IScoreController {
	return &scoreController{}
}

func (sc *scoreController) Answer(c echo.Context) error {
	result := map[string]bool{"result": false}
	username := c.FormValue("username")
	scorelist := c.FormValue("scorelist")
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(scorelist), &data); err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	for _, score := range data {
		add := model.Score{
			Username: username,
			EvalId:   score["eval_id"].(uint),
			Score:    score["score"].(uint),
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
	fileId := c.Param("file_id")
	file_id, _ := strconv.Atoi(fileId)
	evalres, err := sc.eu.GetAllEvals(uint(file_id))
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
