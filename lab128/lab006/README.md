### update&delete

#### 注意点
 - 调用Save之后所生成的sql，会更新所有的属性，就算只更新了一个属性，sql里面也会更新所有
 - Update/Updates可以更新想要更新的属性，updates更新多个属性
 - delete可以用来删除

#### 运行结果
![Imgur](https://i.imgur.com/eWbfnHp.png)

#### 参考资料
 - http://gorm.io/docs/update.html
 - http://gorm.io/docs/delete.html