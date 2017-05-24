### thrift实验2

#### thrift文件
```thrift
namespace go li.rpc
namespace java li.rpc

struct UserDemo{
    1:i32 id;
    2:string name;
    3:i32 age = 15;
    4:string phone;
}

service QuerySrv{
    UserDemo qryUser(1:string name,2:i32 age);

    string queryPhone(1:i32 id);
}
```
namespace后面的名字是用来生成包名的。
service的名字是会在代码里面生成的Service名字

#### 参考资料
 1. https://my.oschina.net/qinerg/blog/165285
 2. http://www.cnblogs.com/cyfonly/p/6059374.html