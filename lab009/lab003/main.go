package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

var dbConfig *DBConfig

func init() {
	readConf()
}

func main() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPwd, dbConfig.DBHost, dbConfig.DBName)
	db, err := sql.Open("mysql", connectInfo)
	if err != nil {
		log.Fatalf("sql.Open error:%v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("db.Ping error:%v", err)
	}
	fmt.Println("数据库连接成功")

	//insert blob类型
	stmtIns, err := db.Prepare("insert into t_test_blob(mailData) value(?)")
	if err != nil {
		log.Fatalf("db.Prepare error:%v", err)
	}
	defer stmtIns.Close()
	data := []byte("helloworld")
	fmt.Println("data=", data)
	_, err = stmtIns.Exec(data)
	if err != nil {
		log.Fatalf("stmtIns.Exec error:%v", err)
	}

	//读
	//rawbytes
	fmt.Println()
	rows, err := db.Query("select * from t_test_blob limit 2")
	if err != nil {
		log.Fatalf("db.Query error:%v", err)
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("rows.Columns error:%v", err)
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	//rows.Scan wants '[]interface{}' as an argument, so we must copy the
	//references into such a slice
	//See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalf("rows.Scan error:%v", err)
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
			fmt.Println("receive data:", col)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("rows.Err error:%v", err)
	}
}

func readConf() {
	//读取数据库配置文件
	data, err := ioutil.ReadFile("../db_config.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	dbConfig = &DBConfig{}
	if err := json.Unmarshal(data, dbConfig); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}
}

type DBConfig struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string
}
