package service

import (
	"gin-Blog/db/mysql"
	"gin-Blog/model"
)

// ExistArticleByID 判断指定ID的article是否存在
func ExistArticleByID(id int) bool {
	var article model.Article
	mysql.MysqlDB.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// GetArticleTotal 根据指定条件获取文章总数
func GetArticleTotal(maps any) int {
	var count int
	mysql.MysqlDB.Model(&model.Article{}).Where(maps).Count(&count)
	return count
}

func GetArticle(id int) (article model.Article) {
	mysql.MysqlDB.Where("id = ?", id).First(&article)
	mysql.MysqlDB.Model(&article).Related(&article.Tag)
	return
}
