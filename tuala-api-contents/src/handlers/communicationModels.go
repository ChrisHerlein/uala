package handlers

import (
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
)

type pageable interface {
	models.Content
}

type page struct {
	Number int `json:"number"`
	Items  any `json:"items"`
}

func toPage[P pageable](number int, items ...P) *page {
	return &page{
		Number: number,
		Items:  items,
	}
}
