package router

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, fc controller.IFileController, sc controller.IScoreController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	e.GET("/getfile", fc.GetFile)
	e.GET("/filereview/:file_id", fc.FileReview)
	e.POST("/fileupload", fc.FileUpload)
	e.GET("/searchfile", fc.SearchFile)
	e.POST("/answer", sc.Answer)
	e.GET("/accessanswer/:file_id", sc.AccessAnswer)
	return e
}
