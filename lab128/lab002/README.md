### 建表操作

#### 注意点
 - 可以在操作之前加上`Debug()`，然后链式调用，这样就可以打印出gorm最终和数据库交互的sql
 - `TableName()`的返回可以当做表名
 - 可以在属性后面加上`gorm:Auto_InCrement`等，这样就会转换成我们自增，或者unique key等，具体的字段可以查看资料

#### 运行结果
![Imgur](https://i.imgur.com/PwUBvdD.png)

#### 参考资料
http://doc.gorm.io/models.html#conventions