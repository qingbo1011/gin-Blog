# gin-Blog
之前已经写过gin的项目了，但是因为一些原因有段时间没有跟进gin了。所以再练手一个gin项目，正好实践一下go1.18。

- [《跟煎鱼学 Go》](https://eddycjy.com/go-categories/)
- [Gin搭建Blog API's](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/) 

在gin-memos项目中使用的是MySQL8，看上面的参考博客，貌似使用的是MySQL5？这里就选用MySQL5.7，顺便测试一下gorm在MySQL5.7上的实践。

对了，这个项目看一半吧，有些用法比较老了。但是里面有些功能，比如**Cron定时任务**、Swagger等等还是值得学习的。

> 这项目我跟着写的时候是有点麻的。毕竟是4年前的项目了，很多处理思路跟我实习和之前处理gin-memos的时候完全不一样。但是还是完全照着敲下来学习学习吧。

## Getting start

在conf目录下创建`config.ini`文件：

```ini
#debug or release
RUN_MODE = debug

[app]
JWT_SECRET = 23347$040412
PAGE_SIZE  = 5

[server]
HTTP_PORT = :8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[mysql]
USER = root
PASSWORD = 123456
HOST = 127.0.0.1
PORT = 3306
NAME = blog
```

## 项目结构

```
gin-Blog
├─.idea
├─conf
├─middleware
├─model
├─pkg
├─route
└─runtime

```

- conf：用于存储配置文件

- middleware：应用中间件
- model：应用数据库模型
- pkg：第三方包
- route：路由逻辑处理
- runtime：应用运行时数据



## 心得

在写代码的过程中，我发现gin的源码其实是很好看很容易看懂的。比如在用到`c.Query`、`c.DefaultQuery`和`c.Param`时，可能还不明确具体是什么意思，如何使用。这个时候点进去看一下源码就可以了：

```go
// Query returns the keyed url query value if it exists,
// otherwise it returns an empty string `("")`.
// It is shortcut for `c.Request.URL.Query().Get(key)`
//     GET /path?id=1234&name=Manu&value=
//        c.Query("id") == "1234"
//        c.Query("name") == "Manu"
//        c.Query("value") == ""
//        c.Query("wtf") == ""
func (c *Context) Query(key string) (value string) {
   value, _ = c.GetQuery(key)
   return
}

/************************************/

// DefaultQuery returns the keyed url query value if it exists,
// otherwise it returns the specified defaultValue string.
// See: Query() and GetQuery() for further information.
//     GET /?name=Manu&lastname=
//     c.DefaultQuery("name", "unknown") == "Manu"
//     c.DefaultQuery("id", "none") == "none"
//     c.DefaultQuery("lastname", "none") == ""
func (c *Context) DefaultQuery(key, defaultValue string) string {
	if value, ok := c.GetQuery(key); ok {
		return value
	}
	return defaultValue
}

/************************************/

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//     router.GET("/user/:id", func(c *gin.Context) {
//         // a GET request to /user/john
//         id := c.Param("id") // id == "john"
//     })
func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}
```













