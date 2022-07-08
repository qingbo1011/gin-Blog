package v1

import (
	"gin-Blog/conf"
	"gin-Blog/pkg/error_data"
	"gin-Blog/service"
	"gin-Blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]any)
	data := make(map[string]any)

	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := error_data.SUCCESS

	data["lists"] = service.GetTags(util.GetPage(c), conf.PageSize, maps)
	data["total"] = service.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error_data.GetMsg(code),
		"data": data,
	})
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
}
