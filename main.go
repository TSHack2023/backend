package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	fileRepository := repository.NewFileRepository(db)
	evalRepository := repository.NewEvalRepository(db)
	scoreRepository := repository.NewScoreRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	fileUsecase := usecase.NewFileUsecase(fileRepository, evalRepository)
	evalUsecase := usecase.NewEvalUsecase(evalRepository)
	scoreUsecase := usecase.NewScoreUsecase(scoreRepository, evalRepository)
	userController := controller.NewUserController(userUsecase)
	fileController := controller.NewFileController(fileUsecase, evalUsecase)
	scoreController := controller.NewScoreController(scoreUsecase, evalUsecase)
	e := router.NewRouter(userController, fileController, scoreController)
	e.Logger.Fatal(e.Start(":8080"))
}
