# 配置服务
数据模型应该是：master来更新数据，其他的worker也随之将数据更新。
先启动master，再启动worker

## NOTICE
 - 这个channel不和lab004中的观察模式一样，观察模式只能使用一次，收到event之后，需要重新创建观察。但是这个channel会一直监控

## 运行结果
### master
![Imgur](https://i.imgur.com/5bujxYe.png)

### worker
![Imgur](https://i.imgur.com/6ZzambL.png)

## 参考资料
 - https://holynull.gitbooks.io/zookeeper/content/pei_zhi_fu_wu_configuration_service.html