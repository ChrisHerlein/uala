package enums

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

var (
	Err400 = errors.New("bad request:")
	Err404 = errors.New("not found:")
	Err409 = errors.New("conflict:")

	Err503 = errors.New("unavailable:")
)

var HttpErrors = map[error]int{
	Err400: fiber.StatusBadRequest,
	Err404: fiber.StatusNotFound,
	Err409: fiber.StatusConflict,
	Err503: fiber.StatusServiceUnavailable,
}

func GetErrorCode(err error, ifNotExists int) int {
	for k, v := range HttpErrors {
		if errors.Is(err, k) {
			return v
		}
	}
	return ifNotExists
}
