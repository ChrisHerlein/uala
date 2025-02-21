package models

import (
	"gorm.io/gorm"

	usersModels "github.com/ChrisHerlein/uala/tuala-api-users/src/models"
)

type Content struct {
	gorm.Model
	AuthorName string           `json:"name"`
	User       usersModels.User `json:"-" gorm:"foreignKey:name;references:AuthorName"`
	Text       string           `json:"text"`
}
