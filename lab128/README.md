# [gorm](https://github.com/jinzhu/gorm)
安装:`go get -u -v github.com/jinzhu/gorm`

# 新版
lab243用的是新版本

|实验|简介|说明|
|---|---|---|
|lab001|连接数据库| |
|lab002|对表的操作，创建表，对column操作，创建索引，gorm的错误处理| |
|lab003|插入，更新，删除| |
|lab004|查询| |
|lab005|数据库连接测试| |
|lab006|测试update操作| |
|lab007|mysql的json数据格式| |
|lab008|事务| |
|lab009|查询为空| |
|lab010|一对多，多对多，联合查询|Association,preload|
|lab011|分库分表||
|lab012|查询一列||

## NOTICE
 - gorm的tag在有auto_increment的情况下，要是还有type，auto_increment会不生效。
 参考资料：https://github.com/jinzhu/gorm/issues/2755