package models

import (
	"github.com/jinzhu/gorm"
)

// Vote permite controlar que un usuario solo
//vote por solo una vez por cada comentario
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentId" gorm:"not null"`
	UserID    uint `json:"userID" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
