# 插入，更新，删除

## NOTICE
1. insert可以用Create，Create的会自动赋值主键
2. NewRecord可以查看主键是否为空
3. 调用Save之后所生成的sql，会更新所有的属性，就算只更新了一个属性，sql里面也会更新所有
4. Update/Updates可以更新想要更新的属性，updates更新多个属性
5. delete可以用来删除

## 参考资料
 - http://gorm.io/docs/create.html
 - http://gorm.io/docs/update.html
 - http://gorm.io/docs/delete.html