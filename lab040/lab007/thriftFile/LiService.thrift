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