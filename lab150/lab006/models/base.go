package models

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

var Db *gorm.DB
var err error
var driverName string
var conn string

/**
*设置数据库连接
*@param diver string
 */
func Register() {
	driverName = "mysql"
	if isTestEnv() { //如果是测试使用测试数据库
		conn = "root:123456@(127.0.0.1:3306)/db_iris?charset=utf8&parseTime=True&loc=Local"
	} else {
		conn = "root:123456@(127.0.0.1:3306)/db_iris?charset=utf8&parseTime=True&loc=Local"
	}
	//初始化数据库
	Db, err = gorm.Open(driverName, conn)
	if err != nil {
		color.Red(fmt.Sprintf("gorm open 错误: %v", err))
	}
}
func IsNotFound(err error) {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); !ok && err != nil {
		color.Red(fmt.Sprintf("error :%v \n ", err))
	}
}

//获取程序运行环境
// 根据程序运行路径后缀判断
//如果是 test 就是测试环境
func isTestEnv() bool {
	files := os.Args
	for _, v := range files {
		if strings.Contains(v, "test") {
			return true
		}
	}
	return false
}

// 接口返回数据对想
type Response struct {
	Status bool        `json:"status"` //接口状态 true ,false
	Msg    interface{} `json:"msg"`    // 接口信息
	Data   interface{} `json:"data"`   //接口数据
}
