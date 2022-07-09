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

// ExistTagByName 判断指定name的Tag是否存在（AddTag时需要先判断）
func ExistTagByName(name string) bool {
	var tag model.Tag
	mysql.MysqlDB.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// AddTag 添加Tag
func AddTag(name string, state int, createdBy string) {
	mysql.MysqlDB.Create(&model.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
}

// ExistTagByID 判断指定id的Tag是否存在（EditTag时需要先判断）
func ExistTagByID(id int) bool {
	var tag model.Tag
	// .Select 指定要从数据库检索的字段(默认情况下，将选择所有字段)
	mysql.MysqlDB.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// EditTag 根据ID修改Tag
func EditTag(id int, data any) {
	// Update(data) 使用map更新多个属性，只会更新那些被更改了的字段
	mysql.MysqlDB.Model(&model.Tag{}).Where("id = ?", id).Update(data)
}
