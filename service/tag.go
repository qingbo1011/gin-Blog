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

func AddTag(name string, state int, createdBy string) {
	mysql.MysqlDB.Create(&model.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
}

// ExistTagByName 判断指定name的Tag是否存在（AddTag时需要先判断）
func ExistTagByName(name string) bool {
	var tag model.Tag
	mysql.MysqlDB.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
