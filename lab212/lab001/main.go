package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 初始化一个 Gorm 适配器并且在一个 Casbin enforcer 中使用它:
	// 这个适配器会使用一个名为 "casbin" 的 MySQL 数据库。
	// 如果数据库不存在，适配器会自动创建它。
	// 你同样也可以像这样 gormadapter.NewAdapterByDB(gormInstance) 使用一个已经存在的 gorm 实例。
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/") //你的驱动和数据源
	if err != nil {
		log.Fatalf("gormadapter.NewAdapter error:%v", err)
	}
	e, _ := casbin.NewEnforcer("./rbac_model.conf", a)

	// 或者你可以像这样使用一个其他的数据库 "abc" :
	// 适配器会使用名为 "casbin_rule" 的数据表。
	// 如果数据表不存在，适配器会自动创建它。
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// 从数据库加载策略规则
	e.LoadPolicy()

	// 检查权限
	if access, _ := e.Enforce("alice", "data1", "read"); access {
		log.Println("access")
	} else {
		log.Println("no access")
	}

	// 更新策略
	// e.AddPolicy()
	// e.RemovePolicy(...)

	e.AddPolicy("alice", "data1", "read")

	// 再次检查权限
	if access, _ := e.Enforce("alice", "data1", "read"); access {
		log.Println("access")
	} else {
		log.Println("no access")
	}

	// 保存策略到数据库
	if err := e.SavePolicy(); err != nil {
		log.Fatalf("savePolicy error:%v", err)
	}
}
