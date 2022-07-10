package v1

import (
	"gin-Blog/model"
	"gin-Blog/pkg/error_data"
	"gin-Blog/service"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := error_data.INVALID_PARAMS
	var data model.Article
	if !valid.HasErrors() {
		if service.ExistArticleByID(id) {
			data = service.GetArticle(id)
			code = error_data.SUCCESS
		} else {
			code = error_data.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error_data.GetMsg(code),
		"data": data,
	})
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context) {

}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
}
