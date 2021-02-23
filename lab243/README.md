# [gorm v2](https://github.com/go-gorm/gorm)
https://gorm.io/zh_CN/docs/v2_release_note.html

## 安装
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

|实验|简介|说明|
|---|---|---|
|lab001|实验新版本，批量插入，动态表名，错误处理| |
|lab002|插件：读写分离| |
|lab003|新版本：更新| |
|lab004|连接postgresql| |

## NOTICE
 - v1版本中的`gorm.IsRecordNotFound`在新版中变为`errors.Is(err, gorm.ErrRecordNotFound)`
 - v1版本中的update在更新之后会去select一下修改的数据，然后修改我们的struct。但是v2现在测试下来，在lab003中看来，v2是不会去select的，但是还是会更新struct的值。
应该是用了map里面的数据直接修改。所以之后update的时候要注意`gorm.Expr`这种更新，因为v2不会去select了，所以这种更新的值拿不到了。