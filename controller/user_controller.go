package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	result := map[string]bool{"result": false}
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	err := uc.uu.SighUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	result["result"] = true
	return c.JSON(http.StatusCreated, result)
}

func (uc *userController) Login(c echo.Context) error {
	result := map[string]bool{"result": false}
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//tokenString, err := uc.uu.Login(user)
	_, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	result["result"] = true
	return c.JSON(http.StatusOK, result)
}

func (uc *userController) LogOut(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
