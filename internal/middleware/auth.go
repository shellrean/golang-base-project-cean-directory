package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/shellrean/golang-base-project-cean-directory/domain"
	"github.com/shellrean/golang-base-project-cean-directory/dto"
	"net/http"
	"strings"
)

func Authenticate(authService domain.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.Split(c.Get("Authorization"), " ")
		if len(token) < 2 {
			return c.Status(http.StatusUnauthorized).JSON(dto.NewResponseMessage("Sorry, the token you entered is invalid. Please check your token and try again or contact customer support for further assistance. Thank you."))
		}
		user, err := authService.Validate(context.Background(), token[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(dto.NewResponseMessage("Sorry, the token you entered is invalid. Please check your token and try again or contact customer support for further assistance. Thank you."))
		}

		c.Locals("x-user", user)
		return c.Next()
	}
}
