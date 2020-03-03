# swagger
在iris里面自动生成swagger

|实验|简介|说明|
|---|---|---|
|lab001|example| |
|lab002|在001的基础上加上swagger的一些配置| |
|lab003|swagger注释| |

## NOTICE
 - 要先运行`swag init`
 - `@Tags accounts`，相当于分组，表示api就在accounts这个分组下
 - ``
 - `@Router /testapi/get-string-by-int/{some_id} [get]`，这个url要加上`@BasePath /api/v1`，才是最终的url

## 参考资料
 - https://github.com/iris-contrib/swagger
 - https://swaggo.github.io/swaggo.io/declarative_comments_format/
