### 创建组，加入组，查看列表，删除组

#### 注意点
 - 创建之前要是已经有Node了，可以用zkCli.sh里面删除
 - `zk.FlagEphemeral`，create时候用这个flag，创建的就会是临时节点，也就是这个session断开的时候，node就会被删除
 - 删除的时候需要版本号，需要版本号也一样才可以删除，但是-1的时候就不用满足这个条件
 - 删除的时候要没有children

#### 运行结果
![Imgur](https://i.imgur.com/jiXKE7G.png)