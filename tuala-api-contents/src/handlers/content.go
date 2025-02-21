package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/services"
)

type Content interface {
	Create() func(fiber.Ctx) error
	Get() func(fiber.Ctx) error
}

type content struct {
	srv services.Content
}

type createContentBodyRequest struct {
	Text string `json:"text"`
}

func (ch *content) Create() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		var body createContentBodyRequest
		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusNotAcceptable)).
				JSON(map[string]string{"error": err.Error()})
		}

		tweet, err := ch.srv.Create(c.Context(), body.Text)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(tweet)
	}
}

// Will retreive content of specific user
func (ch *content) Get() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		userName := c.Params("name")
		pageParam := c.Params("page")
		page, _ := strconv.Atoi(pageParam) // if errored, let retreive first page
		content, err := ch.srv.Read(userName, page)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(toPage(page, content...))
	}
}

func NewContent(service services.Content) *content {
	return &content{
		srv: service,
	}
}
