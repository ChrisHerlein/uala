package main

import (
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
)

type feedPage struct {
	UserName string           `json:"userName"` // user who will receive the content
	Content  []models.Content `json:"content"`
	Order    int              `json:"order"`
}

type control struct {
	UserName   string `json:"userName"`
	MostRecent int    `json:"mostRecent"` // #order of last page
	SizeOfLast int    `json:"sizeOfLast"`
}
