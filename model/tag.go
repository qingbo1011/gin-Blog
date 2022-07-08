package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	State int    `json:"state"`
}
