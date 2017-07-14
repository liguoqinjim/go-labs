package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lab009/lab004/pb"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
)

type DBConf struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string
}

var Conf *DBConf

func init() {
	readConf()
}

func readConf() *DBConf {
	file, err := os.Open("db_config.json")
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("failed to read data:", err)
	}

	dbConf := new(DBConf)
	if err := json.Unmarshal(data, dbConf); err != nil {
		log.Fatalln("failed to parse json:", err)
	}

	Conf = dbConf

	return dbConf
}

func main() {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", Conf.DBUser, Conf.DBPwd, Conf.DBHost, Conf.DBName)
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

	armyGroup := CreateArmyGroup()
	data, err := proto.Marshal(armyGroup)
	if err != nil {
		log.Fatalln("failed to marshal pb:", err)
	}

	//insert blob类型
	stmtIns, err := db.Prepare("insert into t_test_blob_pb(pbData) value(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	fmt.Println("data=", data)
	_, err = stmtIns.Exec(data)
	if err != nil {
		panic(err.Error())
	}

	//读
	//rawbytes
	fmt.Println()
	rows, err := db.Query("select * from t_test_blob_pb limit 2")
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

			if i == 1 { //第二列的时候解析pb
				receiveArmyGroup := new(pb.ArmyGroup)
				if err := proto.Unmarshal(col, receiveArmyGroup); err == nil {
					fmt.Println("receivepb:", receiveArmyGroup)
				} else {
					log.Fatalln("failed to parse:", err)
				}
			}
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
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
