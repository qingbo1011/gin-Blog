# gin-Blog
之前已经写过gin的项目了，但是因为一些原因有段时间没有跟进gin了。所以再练手一个gin项目，正好实践一下go1.18。

- [《跟煎鱼学 Go》](https://eddycjy.com/go-categories/)
- [Gin搭建Blog API's](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/) 

## Getting start

在conf目录下创建`config.ini`文件：

```ini
#debug or release
[running]
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 23347$040412

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = root
PASSWORD = 123456
#127.0.0.1:3306
HOST = 127.0.0.1:3308
NAME = blog
TABLE_PREFIX = blog_
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







