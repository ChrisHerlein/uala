package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/services"
)

// Feed will expose content of all followed users
type Feed interface {
	Recent() func(fiber.Ctx) error
}

type feed struct {
	srv services.Feed
}

func (f *feed) Recent() func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		pageParam := c.Params("page")
		page, _ := strconv.Atoi(pageParam) // if errored, let retreive first page
		feed, err := f.srv.Recent(c.Context(), page)
		if err != nil {
			return c.Status(enums.GetErrorCode(err, fiber.StatusInternalServerError)).
				JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(feed)
	}
}

func NewFeed(service services.Feed) *feed {
	return &feed{
		srv: service,
	}
}
