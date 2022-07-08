package model

import (
	"gin-Blog/db/mysql"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	State int    `json:"state"`
}

// GetTags 分页查询tags
func GetTags(pageNum int, pageSize int, maps any) (tags []Tag) {
	mysql.MysqlDB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagTotal 获取Tag总数
func GetTagTotal(maps any) (count int) {
	mysql.MysqlDB.Model(&Tag{}).Where(maps).Count(&count)
	return
}
