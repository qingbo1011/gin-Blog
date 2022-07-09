package route

import (
	v1 "gin-Blog/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.Default() //gin.Default()和gin.New()的区别：gin.Default()默认使用logger和recovery中间件

	apiV1 := r.Group("/api/v1")
	{
		// Tag相关
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag) // gin绑定URI
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		// Article相关
		apiV1.GET("/articles", v1.GetArticles)          //获取文章列表
		apiV1.GET("/articles/:id", v1.GetArticle)       //获取指定文章
		apiV1.POST("/articles", v1.AddArticle)          //新建文章
		apiV1.PUT("/articles/:id", v1.EditArticle)      //更新指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle) //删除指定文章
	}

	return r
}
