### 获取服务，获取一组值

#### 注意点
 - `rsp, err := cli.Get(context.TODO(), KEY_NAME_PREFIX+"/", clientv3.WithPrefix())`可以得到一组值

#### 运行结果
![Imgur](https://i.imgur.com/uOk8IbR.png)