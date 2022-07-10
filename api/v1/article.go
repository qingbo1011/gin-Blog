package v1

import (
	"gin-Blog/conf"
	"gin-Blog/model"
	"gin-Blog/pkg/error_data"
	"gin-Blog/service"
	"gin-Blog/util"
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
	data := make(map[string]any)
	maps := make(map[string]any)

	valid := validation.Validation{}
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	tagID := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagID
		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := error_data.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error_data.SUCCESS
		data["lists"] = service.GetArticles(util.GetPage(c), conf.PageSize, maps)
		data["total"] = service.GetArticleTotal(maps)
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

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	tagID := com.StrTo(c.Query("tag_id")).MustInt()            // 文章标签
	title := c.Query("title")                                  // 文章标题
	desc := c.Query("desc")                                    // 文章简述
	content := c.Query("content")                              // 文章内容
	createdBy := c.Query("created_by")                         // 文章创建人
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt() // 文章状态

	valid := validation.Validation{}
	valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := error_data.INVALID_PARAMS
	if !valid.HasErrors() {
		if service.ExistTagByID(tagID) {
			code = error_data.SUCCESS
			data := make(map[string]any)
			data["tag_id"] = tagID
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			service.AddArticle(data)
		} else {
			code = error_data.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error_data.GetMsg(code),
	})

}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := error_data.INVALID_PARAMS
	if !valid.HasErrors() {
		if service.ExistArticleByID(id) {
			if service.ExistTagByID(tagId) {
				code = error_data.SUCCESS
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy

				service.EditArticle(id, data)
			} else {
				code = error_data.ERROR_NOT_EXIST_TAG
			}
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
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := error_data.INVALID_PARAMS
	if !valid.HasErrors() {
		if service.ExistArticleByID(id) {
			code = error_data.SUCCESS
			service.DeleteArticle(id)
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
	})
}
