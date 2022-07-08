package service

import (
	"gin-Blog/db/mysql"
	"gin-Blog/model"
)

// GetTags 分页查询tags
func GetTags(pageNum int, pageSize int, maps any) (tags []model.Tag) {
	mysql.MysqlDB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagTotal 获取Tag总数
func GetTagTotal(maps any) (count int) {
	mysql.MysqlDB.Model(&model.Tag{}).Where(maps).Count(&count)
	return
}
