package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type DBConfig struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string
}

func getDBConfig(data []byte) *DBConfig {
	var config DBConfig
	json.Unmarshal(data, &config)
	return &config
}

var dbConfig DBConfig

func main() {
	//读取数据库配置文件
	data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig = *getDBConfig(data)

	//连接数据库
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPwd, dbConfig.DBHost, dbConfig.DBName)
	db, err := sql.Open("mysql", connectInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功")

	//insert
	stmtIns, err := db.Prepare("insert into squareNum values(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
	for i := 0; i < 25; i++ {
		_, err := stmtIns.Exec(i, i*i)
		if err != nil {
			panic(err.Error())
		}
	}

	//select
	stmtOut, err := db.Prepare("select squareNumber from squareNum where number = ?")
	if err != nil {
		panic(err.Error())
	}
	var squareNum int
	err = stmtOut.QueryRow(13).Scan(&squareNum)
	fmt.Printf("squareNumber = %d\n", squareNum)
	err = stmtOut.QueryRow(1).Scan(&squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("squareNumber = %d\n", squareNum)

	//rawbytes
	fmt.Println()
	rows, err := db.Query("select * from squareNum limit 2")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
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
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
