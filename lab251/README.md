# [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

|实验|简介|说明|
|---|---|---|
|lab001|demo|链接，查询 |

## 参考资料
 - [文档](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#section-documentation)
 - https://www.mongodb.com/golang
 - https://docs.mongodb.com/manual/reference/operator/query/regex/

## NOTICE
 - 正则和前缀(mongodb里面加了^可以表达前缀)，`r, err = collection.DeleteMany(context.TODO(), bson.D{{"id", primitive.Regex{Pattern: fmt.Sprintf("^%d_*_*", id), Options: ""}}})`