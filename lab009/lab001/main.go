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
	//连接数据库
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

	//insert
	stmtIns, err := db.Prepare("insert into squareNum values(?,?)")
	if err != nil {
		log.Fatalf("db.Prepare error:%v", err)
	}
	defer stmtIns.Close()
	for i := 0; i < 25; i++ {
		_, err := stmtIns.Exec(i, i*i)
		if err != nil {
			log.Fatalf("stmtIns.Exec error:%v", err)
		}
	}

	//select
	stmtOut, err := db.Prepare("select squareNumber from squareNum where number = ?")
	if err != nil {
		log.Fatalf("db.Prepare error:%v", err)
	}
	var squareNum int
	err = stmtOut.QueryRow(13).Scan(&squareNum)
	fmt.Printf("squareNumber = %d\n", squareNum)
	err = stmtOut.QueryRow(1).Scan(&squareNum)
	if err != nil {
		log.Fatalf("stmtOut.QueryRow error:%v", err)
	}
	fmt.Printf("squareNumber = %d\n", squareNum)

	//rawbytes
	fmt.Println()
	rows, err := db.Query("select * from squareNum limit 2")
	if err != nil {
		log.Fatalf("db.Query error:%v", err)
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("rows.Columns error:%v", err)
	} else {
		log.Println("columns:", columns)
	}

	values := make([]sql.RawBytes, len(columns))
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
