package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(255);not null"`
	Author string `json:"author" gorm:"type:varchar(255);not null"`
	Body string `json:"body" gorm:"type:text;not null"`
}
