package route

import "github.com/gin-gonic/gin"

func NewRoute() *gin.Engine {
	r := gin.Default() //gin.Default()和gin.New()的区别：gin.Default()默认使用logger和recovery中间件

	return r
}
