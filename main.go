package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shellrean/golang-base-project-cean-directory/internal/api"
	"github.com/shellrean/golang-base-project-cean-directory/internal/config"
	"github.com/shellrean/golang-base-project-cean-directory/internal/connection"
	"github.com/shellrean/golang-base-project-cean-directory/internal/repository"
	"github.com/shellrean/golang-base-project-cean-directory/internal/service"
)

func main() {
	cnf := config.Get()

	dbConnection := connection.GetDatabase(cnf.Database)

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewAuth(cnf, userRepository)

	app := fiber.New()

	api.NewAuth(app, authService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
