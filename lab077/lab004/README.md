### 读写锁

#### 注意
 - 当写锁阻塞时，新的读锁是无法申请的，这可以有效防止写者饥饿。
 - 读写锁简单来说就是可以由任意数量的读者同时使用，或者只由一个写者使用的锁。(可以有多个同时读)

#### 运行结果
![Imgur](https://i.imgur.com/DArvnpm.png)

#### 参考资料
http://zablog.me/2017/09/27/go_sync/