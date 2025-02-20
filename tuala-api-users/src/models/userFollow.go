package models

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserRefer uint `gorm:"uniqueIndex:idx_user_follows"`
	Follows   uint `gorm:"uniqueIndex:idx_user_follows"`
}
