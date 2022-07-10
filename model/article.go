package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"` // gorm:index，声明这个字段为索引
	Tag   Tag `json:"tag"`                 // 嵌套的struct，利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
