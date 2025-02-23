package main

import (
	"time"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
)

// Will be stored into Redis with key: `{userId}-{page.Order}`
type feedPage struct {
	UserName string           `json:"userName"` // user who will receive the content
	Content  []models.Content `json:"content"`  // Newest should be at begining of slice
	Order    int              `json:"order"`
}

// Will be stored into Redis with key: `{userId}-control`
type control struct {
	UserName   string    `json:"userName"`
	MostRecent int       `json:"mostRecent"` // #order of last page
	SizeOfLast int       `json:"sizeOfLast"` // how many tweets are stored into last page
	Last       time.Time `json:"last"`       // Refers to the newest content in last page
}
