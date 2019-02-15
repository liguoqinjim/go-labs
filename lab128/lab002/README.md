# 对表的操作，创建表，新增字段，创建索引，gorm的错误处理

## NOTICE
1. 结构体里面的字段名的小写如果是id的话，gorm会默认这个为主键
2. 可以在结构体里面加上`gorm.Model`，那么自动生成的数据库表会加上4个字段
3. 可以在操作之前加上`Debug()`，然后链式调用，这样就可以打印出gorm最终和数据库交互的sql
4. `TableName()`的返回可以当做表名
5. 可以在属性后面加上`gorm:Auto_InCrement`等，这样就会转换成我们自增，或者unique key等，具体的字段可以查看资料
6. DefaultTableNameHandler，可以修改默认的表名规则
7. golang的`time.Time`会对应sql里面的timestamp
8. `gorm:"column:beast_id"`，这样可以指定列名

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
  `id` varchar(255) NOT NULL DEFAULT '',
  `uid` int(11) DEFAULT NULL,
  `uname` varchar(255) DEFAULT NULL,
  `uage` int(11) DEFAULT NULL,
  `stu_id` int(11) DEFAULT NULL,
  `udes` varchar(255) DEFAULT NULL,
  `u_address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type Student struct {
	Id    int
	Sno   int
	Sname string
	Sage  int
}
```
```mysql
CREATE TABLE `students` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sno` int(11) DEFAULT NULL,
  `sname` varchar(255) DEFAULT NULL,
  `sage` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type Class struct {
	Id  int `gorm:"AUTO_INCREMENT"`
	Cno string
}
```
```mysql
CREATE TABLE `t_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cno` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type User2 struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard // One-To-One relationship (has one - use CreditCard's UserID as foreign key)
	Emails     []Email    // One-To-Many relationship (has many - use Email's UserID as foreign key)

	BillingAddress   Address // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // Ignore this field
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}
```
```mysql
CREATE TABLE `t_user2` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `birthday` timestamp NULL DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `num` int(11) DEFAULT NULL,
  `billing_address_id` bigint(20) DEFAULT NULL,
  `shipping_address_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_t_user2_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // Foreign key (belongs to), tag `index` will create index for this column
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	Subscribed bool
}
```
```mysql
CREATE TABLE `t_emails` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `subscribed` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_t_emails_email` (`email`),
  KEY `idx_t_emails_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}
```
```mysql
CREATE TABLE `t_addresses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `address1` varchar(255) NOT NULL,
  `address2` varchar(100) DEFAULT NULL,
  `post` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `address1` (`address1`),
  UNIQUE KEY `address2` (`address2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}
```
```mysql
CREATE TABLE `t_languages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_name_code` (`name`,`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

```golang
type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
```
```mysql
CREATE TABLE `t_credit_cards` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `number` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_t_credit_cards_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

<hr>

many2many生成的表
```mysql
CREATE TABLE `t_user_languages` (
  `user2_id` int(10) unsigned NOT NULL DEFAULT '0',
  `language_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user2_id`,`language_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

## 参考资料
 - http://doc.gorm.io/models.html#conventions
 - http://gorm.io/docs/error_handling.html
