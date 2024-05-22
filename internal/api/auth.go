package api

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shellrean/golang-base-project-cean-directory/domain"
	"github.com/shellrean/golang-base-project-cean-directory/dto"
	"net/http"
	"time"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, authHandler fiber.Handler,
	authService domain.AuthService) {

	ha := authApi{
		authService: authService,
	}

	app.Post("/v1/authenticate", ha.authenticate)
	app.Post("/v1/authenticate/validate", authHandler, ha.authenticateValidate)
}

func (a authApi) authenticate(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, err := a.authService.Authenticate(c, req)

	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredential) {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.NewResponseMessage("Invalid credential. Please check your username and password, then try again. If the problem persists, contact support for assistance."))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("An internal server error has occurred. Please try again later. If the issue persists, contact support for further assistance. We apologize for any inconvenience."))
	}

	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.AuthRes](res))
}

func (a authApi) authenticateValidate(ctx *fiber.Ctx) error {
	userLocal := ctx.Locals("x-user")
	if userLocal == nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("Sorry, the token you entered is invalid. Please check your token and try again or contact customer support for further assistance. Thank you."))
	}
	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.UserData](userLocal.(dto.UserData)))
}
