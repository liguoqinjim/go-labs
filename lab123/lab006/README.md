# 设置cookie，设置不一样的domain
我们要在打开一个页面之前，给这个页面设置cookie。

## 注意
 - 如果打开这个页面之后，再设置cookie，cookie的值会不能发给服务器
 - page没打开页面之前，直接调用设置setCookie的方法，虽然不会报错，但是之后查看控制台也是看不到这个cookie的
 - 可以和这个程序里面一样，先打开example.com，然后设置另一个界面的cookie。然后再打开另一个界面的时候，就可以看到这个cookie了