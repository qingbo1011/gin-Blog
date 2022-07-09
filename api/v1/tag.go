package v1

import (
	"gin-Blog/conf"
	"gin-Blog/pkg/error_data"
	"gin-Blog/service"
	"gin-Blog/util"
	"net/http"

	"github.com/astaxie/beego/validation"
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
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("created_by")

	// 表单验证时用到validation包
	// 可以发现在beego的validation中，限制参数类型是需要指定的。在gin中，ShouldBind绑定的struct可以通过添加tag来实现（参考gin-memos项目）
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 50, "name").Message("名称最长为50字符")
	valid.Required(createBy, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := error_data.INVALID_PARAMS
	if !valid.HasErrors() {
		if !service.ExistTagByName(name) {
			code = error_data.SUCCESS
			service.AddTag(name, state, createBy)
		} else {
			code = error_data.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error_data.GetMsg(code),
	})

}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
}
