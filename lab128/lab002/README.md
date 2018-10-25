# 对表的操作，创建表，新增字段，创建索引

## NOTICE
1. 结构体里面的字段名的小写如果是id的话，gorm会默认这个为主键
2. 可以在结构体里面加上`gorm.Model`，那么自动生成的数据库表会加上4个字段
3. 可以在操作之前加上`Debug()`，然后链式调用，这样就可以打印出gorm最终和数据库交互的sql
4. `TableName()`的返回可以当做表名
5. 可以在属性后面加上`gorm:Auto_InCrement`等，这样就会转换成我们自增，或者unique key等，具体的字段可以查看资料
6. DefaultTableNameHandler，可以修改默认的表名规则

## struct和对应的表结构
```golang
type User struct {
	ID       string
	Uid      int
	Uname    string
	Uage     int
	StuId    int
	Udes     string
	UAddress string
}
```
```mysql
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uid` int(11) DEFAULT NULL,
  `uname` varchar(255) DEFAULT NULL,
  `uage` int(11) DEFAULT NULL,
  `stu_id` int(11) DEFAULT NULL,
  `udes` varchar(255) DEFAULT NULL,
  `u_address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

```

## 参考资料
http://doc.gorm.io/models.html#conventions