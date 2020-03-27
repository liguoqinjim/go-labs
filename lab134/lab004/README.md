# 观察模式

## NOTICE
 - 节点删除或者节点创建，观察模式都会收到事件的
 - 观察模式只能用一次，在接收到event之后，就要重新获得观察模式才会获得新的event。这个特性是zookeeper决定的。![Imgur](https://imgur.com/undefined)

## 参考资料
 - https://holynull.gitbooks.io/zookeeper/content/shu_ju_mo_xing_data_model.html

## 运行结果
![Imgur](https://i.imgur.com/4yZ8viP.png)