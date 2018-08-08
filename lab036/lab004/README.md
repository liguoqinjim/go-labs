# 打印出方法名和行号01，demo
logrus本身没有方法可以打印行号，这里要使用hook才行。目前看到的是可以用参考资料1里面的hook

## 注意点
 - 我们可以控制hook的skip参数，也就是hook里面调用runtime.caller的参数
 - hook里面的逻辑是一旦runtime.caller的file不是logrus，就退出逻辑了。如果有需要的话，是要自己改源码的

## 运行结果
![Imgur](https://i.imgur.com/wk1NS7M.png)

## 参考资料
1. https://github.com/onrik/logrus
2. https://github.com/sirupsen/logrus/issues/63