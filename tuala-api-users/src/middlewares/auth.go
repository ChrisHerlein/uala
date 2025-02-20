package middlewares

import (
	"context"
	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/repositories"
)

func Auth(db repositories.DB) func(c fiber.Ctx) error {
	var setter = func(c fiber.Ctx) error {
		ctx := c.Context()
		headers := c.GetReqHeaders()

		userName := headers[enums.HeaderUserName]
		if len(userName) == 0 {
			return c.Status(fiber.StatusUnauthorized).
				JSON(map[string]string{"error": "no user name provided"})
		}

		password := headers[enums.HeaderUserPassword]
		if len(password) == 0 {
			return c.Status(fiber.StatusUnauthorized).
				JSON(map[string]string{"error": "no password provided"})
		}

		user, err := db.Get(userName[0], password[0])
		if err != nil {
			return c.Status(fiber.StatusForbidden).
				JSON(map[string]string{"error": "invalid credentials"})
		}

		ctx = context.WithValue(ctx, enums.CtxUserID, user.ID)
		ctx = context.WithValue(ctx, enums.CtxUserName, userName[0])

		c.SetContext(ctx)
		return c.Next()
	}
	return setter
}
