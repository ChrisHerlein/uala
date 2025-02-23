package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name"`
	Password string   `json:"-"` // bad idea to send it to other users
	Follows  []Follow `json:"follows" gorm:"foreignKey:UserRefer;references:id;constraint:OnDelete:CASCADE"`
}
