package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

func main() {
	//读取配置文件
	data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		log.Fatal("readfile error ", err)
	}

	//解析为DBConfig
	dbConfig := getDBConfig(data)
	if dbConfig == nil {
		log.Fatal("getDBConfig error ", dbConfig)
	}

	//连接数据库
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPwd, dbConfig.DBHost, dbConfig.DBName)
	db, err := sql.Open("mysql", connectInfo)
	if err != nil {
		log.Fatal("sql.Open error ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("db.ping error ", err)
	}
	fmt.Println("数据库连接成功")

	//insert
	stmtIns, err := db.Prepare("insert into student(sid,sname,sage) values(?,?,?)")
	defer stmtIns.Close()
	if err != nil {
		log.Fatal("stmtIns error ", err)
	}
	var lastId int
	for i := 1; i <= 3; i++ {
		result, err := stmtIns.Exec(i, fmt.Sprintf("小明%d", i), 20+i)
		if err != nil {
			log.Fatal("stmtIns insert error ", err)
		}
		newId, err := result.LastInsertId() //新插入的自增id
		if err != nil {
			log.Fatal("newId error ", err)
		}
		fmt.Println("insert id ", newId)
		lastId = int(newId)
	}

	//查找一行
	stmtOut, err := db.Prepare("select * from student where id = ?")
	if err != nil {
		log.Fatal("stmtOut prepare error ", err)
	}
	rowResult := make([]interface{}, 4)
	rowValues := make([]string, 4)
	for i := range rowResult {
		rowResult[i] = &rowValues[i]
	}
	err = stmtOut.QueryRow(lastId).Scan(rowResult...)
	if err != nil {
		log.Fatal("stmtOut queryRow error ", err)
	}
	fmt.Println(rowValues)

	//查询前三行
	rows, err := db.Query("select * from student limit 3")
	if err != nil {
		log.Fatal("db.Query error ", err)
	}
	coloums, err := rows.Columns()
	scanArgs := make([]interface{}, len(coloums))
	values := make([]sql.RawBytes, len(coloums))
	for i := range scanArgs {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal("rows scan error ", err)
		}

		var value string
		for n, v := range values {
			if v == nil {
				value = "NULL"
			} else {
				value = string(v)
			}
			fmt.Printf("%s:%s ", coloums[n], value)
		}
		fmt.Println()
	}
}
