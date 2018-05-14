### watch
用etcdctl对值进行修改，然后查看watch模式收到的event

#### 注意点
 - 不像zookeeper里面，观察模式只能用一次，下一次就要用新的。
 etcd里面的watch可以一直使用的
 - 对值的修改、删除，watch都会收到event

#### 运行结果
##### 运行结果
![Imgur](https://i.imgur.com/FDayKTi.png)

##### etcdctl
![Imgur](https://i.imgur.com/SRlh0Tp.png)