package types

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Username uint
	Message  string `json:"message"`
}
