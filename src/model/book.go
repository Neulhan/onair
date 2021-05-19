package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" validate:"required,min=3,max=32"`
	Author string `json:"author" validate:"required,min=3,max=32"`
	Rating int    `json:"rating" validate:"required"`
}

func (Book) TableName() string {
	return "books"
}
