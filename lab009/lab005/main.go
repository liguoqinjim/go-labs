package main

import (
	"database/sql"
	b64 "encoding/base64"
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
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功")

	//insert blob类型
	//stmtIns, err := db.Prepare("insert into t_test_blob(mailData) value(?)")
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer stmtIns.Close()
	//data := []byte("helloworld")
	//fmt.Println("data=", data)
	//_, err = stmtIns.Exec(data)
	//if err != nil {
	//	panic(err.Error())
	//}

	//读
	//rawbytes
	fmt.Println()
	rows, err := db.Query("select * from t_mail where id = '890851897148182528'")
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

	//rows.Scan wants '[]interface{}' as an argument, so we must copy the
	//references into such a slice
	//See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	mail := &Mail{}

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
			fmt.Println("receive data:", col)
			fmt.Println("receive data length:", len(col))
			if i == 0 {
				mail.MailId = value
			} else if i == 6 {
				mail.MailMessage = value
			} else if i == 10 {
				mail.MailData = value
			}
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//测试转到json的时候的问题
	fmt.Println("测试转json")
	fmt.Printf("%+v\n", mail)
	fmt.Println(len(mail.MailId), len(mail.MailMessage), len(mail.MailData))

	data, err := json.Marshal(mail)
	if err != nil {
		log.Fatalln("json.Marshal error:", err)
	}
	fmt.Println(string(data))

	mail2 := new(Mail)
	err2 := json.Unmarshal(data, mail2)
	if err2 != nil {
		log.Fatalln("json.Unmarshal error:", err)
	}
	fmt.Printf("%+v\n", mail2)
	fmt.Println(len(mail2.MailId), len(mail2.MailMessage), len(mail2.MailData))

	//测试base64转码
	mail3 := &Mail{MailId: mail.MailId, MailMessage: mail.MailMessage, MailData: mail.MailData}
	//mailData转码到base64
	mailDataBase64 := b64.URLEncoding.EncodeToString([]byte(mail.MailData))
	mail3.MailData = mailDataBase64
	data3, err := json.Marshal(mail3)
	if err != nil {
		log.Fatalln("json.Marshal error:", err)
	}
	fmt.Printf("data3=%s\n", data3)
	fmt.Println(len(mail3.MailId), len(mail3.MailMessage), len(mail3.MailData))

	//base64->string
	mail4 := new(Mail)
	err3 := json.Unmarshal(data3, mail4)
	if err3 != nil {
		log.Fatalln("json.Unmarshal error:", err)
	}
	fmt.Printf("mail4=%+v\n", mail4)
	fmt.Println(len(mail4.MailId), len(mail4.MailMessage), len(mail4.MailData))
	mailDataBase64Decode, _ := b64.URLEncoding.DecodeString(mail4.MailData)
	fmt.Println("mailDataBase64Decode:", mailDataBase64Decode)
	mail4.MailData = string(mailDataBase64Decode)
	fmt.Println("mail4.MailData.len=", len(mail4.MailData))

	//比较mail里面的mailData和mail4的mailData
	fmt.Println("比较mail和mail4")
	fmt.Println([]byte(mail.MailData))
	fmt.Println([]byte(mail4.MailData))
}

type Mail struct {
	MailId      string
	MailMessage string
	MailData    string
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
