package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lab009/lab004/pb"
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

	armyGroup := CreateArmyGroup()
	data, err := proto.Marshal(armyGroup)
	if err != nil {
		log.Fatalf("failed to marshal pb:%v", err)
	}

	//insert blob类型
	stmtIns, err := db.Prepare("insert into t_test_blob_pb(pbData) value(?)")
	if err != nil {
		log.Fatalf("db.Prepare error:%v", err)
	}
	defer stmtIns.Close()

	fmt.Println("data=", data)
	_, err = stmtIns.Exec(data)
	if err != nil {
		log.Fatalf("stmtIns.Exec error:%v", err)
	}

	//读
	//rawbytes
	fmt.Println()
	rows, err := db.Query("select * from t_test_blob_pb limit 2")
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

			if i == 1 { //第二列的时候解析pb
				receiveArmyGroup := new(pb.ArmyGroup)
				if err := proto.Unmarshal(col, receiveArmyGroup); err == nil {
					fmt.Println("receivepb:", receiveArmyGroup)
				} else {
					log.Fatalf("failed to parse:%v", err)
				}
			}
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("rows.Err error:%v", err)
	}
}

//ArmyGroup
func CreateArmyGroup() *pb.ArmyGroup {
	h1 := &pb.Hero{ID: "H001", HeroID: 1, HeroLv: 2, HeroStar: 3}
	h2 := &pb.Hero{ID: "H002", HeroID: 11, HeroLv: 12, HeroStar: 13}

	s1 := &pb.Soldier{ID: "S001", SoldierId: 101, SoldierNum: 500}
	s2 := &pb.Soldier{ID: "S002", SoldierId: 102, SoldierNum: 1000}

	armyGroup := &pb.ArmyGroup{ID: "A001", Heros: []*pb.Hero{h1, h2}, Soldiers: []*pb.Soldier{s1, s2}}
	return armyGroup
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
