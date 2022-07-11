# gin-Blog
之前已经写过gin的项目了，但是因为一些原因有段时间没有跟进gin了。所以再练手一个gin项目，正好实践一下go1.18。

- [《跟煎鱼学 Go》](https://eddycjy.com/go-categories/)
- [Gin搭建Blog API's](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/) 

在gin-memos项目中使用的是MySQL8，看上面的参考博客，貌似使用的是MySQL5？这里就选用MySQL5.7，顺便测试一下gorm在MySQL5.7上的实践。

对了，这个项目看一半吧，有些用法比较老了。但是里面有些功能，比如**Cron定时任务**等等还是值得学习的。

> 这项目我跟着写的时候是有点麻的。毕竟是4年前的项目了，很多处理思路跟我实习和之前处理gin-memos的时候完全不一样。但是还是完全照着敲下来学习学习吧。

## 实现功能

- Gin的基本路由route和API（GET、POST、PUT、DELETE）
- Gin整合Gorm实现对MySQL的增删改查
- Cron定时任务

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

## JWT

JWT的学习并没有参考这篇博客。因为他直接把token作为GET Param传进来了，而一般我们都是放在header的。可以参考我的另一个学习项目。（[Gin+JWT+Air热部署的简单实践](https://www.qingbo1011.top/2022/05/01/Gin+JWT+Air%E7%83%AD%E9%83%A8%E7%BD%B2%E7%9A%84%E7%AE%80%E5%8D%95%E5%AE%9E%E8%B7%B5/)）

## Cron定时任务

**[Go 每日一库之 cron](https://segmentfault.com/a/1190000023029219)**

[掘金：Golang 定时任务cron最新版](https://juejin.cn/post/7004656484902502408)

 [cron](https://github.com/robfig/cron) 包，实现了`cron`规范解析器和任务运行器，简单来讲就是包含了定时任务所需的功能。

### Cron 表达式格式

|             字段名             | 是否必填 | 允许的值        | 允许的特殊字符 |
| :----------------------------: | -------- | --------------- | -------------- |
|         秒（Seconds）          | Yes      | 0-59            | * / , -        |
|         分（Minutes）          | Yes      | 0-59            | * / , -        |
|          时（Hours）           | Yes      | 0-23            | * / , -        |
| 一个月中的某天（Day of month） | Yes      | 1-31            | * / , - ?      |
|          月（Month）           | Yes      | 1-12 or JAN-DEC | * / , -        |
|     星期几（Day of week）      | Yes      | 0-6 or SUN-SAT  | * / , - ?      |

- Cron表达式表示一组时间，使用6个空格分隔的字段
- 可以留意到Golang的Cron比Crontab多了一个秒级，以后遇到秒级要求的时候就省事了

### Cron 特殊字符

- 星号 ( `*` )：将匹配字段的所有值；
- 斜线 ( `/` )：描述范围的增量，表现为 `N-MAX/x`，`first-last/x` 的形式，例如 `3-59/15` 表示此时的**第3分钟和此后的每15分钟，到59分钟为止**。即从 N 开始，使用增量直到该特定范围结束。它不会重复；
- 逗号 ( `,` )：用于分隔列表中的项目。例如，在 Day of week 使用`MON,WED,FRI`将意味着星期一，星期三和星期五；
- 连字符 ( `-` )：用于定义范围。例如，`9-17` 表示从上午 9 点到下午 5 点的每个小时；
- 问号 ( `?` )：不指定值，用于代替 `*`，类似`_` 的存在，不难理解。

### 预定义的 Cron 时间表

| 输入                       | 简述                                   | 相当于      |
| -------------------------- | -------------------------------------- | ----------- |
| `@yearly` (or `@annually`) | 1 月 1 日午夜运行一次                  | 0 0 0 1 1 * |
| `@monthly`                 | 每个月的午夜，每个月的第一个月运行一次 | 0 0 0 1 * * |
| `@weekly`                  | 每周一次，周日午夜运行一次             | 0 0 0 * * 0 |
| `@daily` (or `@midnight`)  | 每天午夜运行一次                       | 0 0 0 * * * |
| `@hourly`                  | 每小时运行一次                         | 0 0 * * * * |

> 具体实践参考代码。





## 心得

1.在写代码的过程中，我发现gin的源码其实是很好看很容易看懂的。比如在用到`c.Query`、`c.DefaultQuery`和`c.Param`时，可能还不明确具体是什么意思，如何使用。这个时候点进去看一下源码就可以了：

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

2.这个系列的教程很多地方都是用`map[string]any`类型，我觉得使用`struct`来替换`map`的使用，会方便很多。

3为了方便展示原博客作者选用了 `GET/Param` 传参的方式，而在gin-memos项目中使用的大多则是在`Request Body`中传入`json`的。（`GET/Param` 传参的方式肯定是不行的，因为url长度是有限的，文章存储的内容就有限了。）



