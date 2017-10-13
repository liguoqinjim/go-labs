### 实验cookie

#### 注意点
1. 注意看request2调用url的时候，cookie的变化。gorequest库，在每次访问新的链接的时候都会clear掉之前的状态。具体可以看参考资料1
2. 有redirect的时候，可以用`RedirectPolicy`

#### 参考资料
1. https://github.com/parnurzeal/gorequest/issues/76

#### 运行结果
![Imgur](https://i.imgur.com/M4y9Bji.png)