package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/services"
)

type Users interface {
	Create() func(fiber.Ctx) error
	Get() func(fiber.Ctx) error
	Follow() func(fiber.Ctx) error
	Unfollow() func(fiber.Ctx) error
}

type users struct {
	srv services.Users
}

type createUserBodyRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *users) Create() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var body createUserBodyRequest
		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusNotAcceptable)).
				JSON(map[string]string{"error": err.Error()})
		}

		user, err := u.srv.Create(body.Name, body.Password)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

func (u *users) Get() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		userName := c.Params("name")
		user, err := u.srv.Get(userName)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

func (u *users) Follow() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		userName := c.Params("name")
		user, err := u.srv.Follow(c.Context(), userName)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func (u *users) Unfollow() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		userName := c.Params("name")
		user, err := u.srv.Unfollow(c.Context(), userName)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func NewUsers(service services.Users) *users {
	return &users{
		srv: service,
	}
}
