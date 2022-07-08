package route

import (
	v1 "gin-Blog/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.Default() //gin.Default()和gin.New()的区别：gin.Default()默认使用logger和recovery中间件

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag) // gin绑定URI
		apiV1.GET("/tags/:id", v1.DeleteTag)
	}

	return r
}
