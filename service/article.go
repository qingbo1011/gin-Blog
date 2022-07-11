package service

import (
	"fmt"
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

// GetArticleTotal 根据指定条件获取article总数
func GetArticleTotal(maps any) int {
	var count int
	mysql.MysqlDB.Model(&model.Article{}).Where(maps).Count(&count)
	return count
}

// GetArticle 根据id获取单个article
func GetArticle(id int) (article model.Article) {
	mysql.MysqlDB.Where("id = ?", id).First(&article)
	// 这里是在关联tag表。能够达到关联，首先是gorm本身做了大量的约定俗成：
	// Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID的方式去找到这两个类之间的关联关系
	// Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
	mysql.MysqlDB.Model(&article).Related(&article.Tag)
	return
}

// GetArticles 根据条件maps，返回符合条件的所有article
func GetArticles(pageNum int, pageSize int, maps any) (articles []model.Article) {
	// 这里使用Preload，是为了查询每一项article的关联tag。Preload就是一个预加载器，它会执行两条SQL：
	// SELECT * FROM article; 和 SELECT * FROM tag WHERE id IN (1,2,3,4...);
	// 在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到article的tag中，会特别方便，并且避免了循环查询
	// 实现这种功能的其他两种方法：1.gorm的Join 2.循环Related（综合之下，还是Preload更好）
	mysql.MysqlDB.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// AddArticle 新增article
func AddArticle(data map[string]any) {
	mysql.MysqlDB.Create(&model.Article{
		// 通过断言将any类型转为我们需要的类型
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
}

// EditArticle 修改article
func EditArticle(id int, data any) {
	mysql.MysqlDB.Model(&model.Article{}).Where("id = ?", id).Update(data)
}

// DeleteArticle 根据id删除article
func DeleteArticle(id int) {
	mysql.MysqlDB.Where("id = ?", id).Delete(&model.Article{})
}

// CleanAllArticle 硬删除article表中所有已经被(软)删除的数据
func CleanAllArticle() error {
	err := mysql.MysqlDB.Unscoped().Where("deleted_at != ?", "null").Delete(&model.Article{}).Error
	return err
}

// CronArticle 测试Corn定时任务用
func CronArticle() {
	fmt.Println("关于Article的定时任务")
}
