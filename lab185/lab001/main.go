package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	dbConfig := readConf()
	log.Println("dbConfig=", dbConfig)

	return

	//cfg := canal.NewDefaultConfig()
	//cfg.Addr = "hp-112:3306"
	//cfg.User = "root"
	//
	//// We only care table canal_test in test db
	//cfg.Dump.TableDB = "test"
	//cfg.Dump.Tables = []string{"canal_test"}
	//
	//c, err := NewCanal(cfg)
	//
	//type MyEventHandler struct {
	//	DummyEventHandler
	//}
	//
	//func(h *MyEventHandler) OnRow(e * RowsEvent)
	//error{
	//	log.Infof("%s %v\n", e.Action, e.Rows)
	//	return nil
	//}
	//
	//func(h *MyEventHandler) String()
	//string{
	//	return "MyEventHandler"
	//}
	//
	//// Register a handler to handle RowsEvent
	//c.SetEventHandler(&MyEventHandler{})
	//
	//// Start canal
	//c.Start()
}

func readConf() *DBConfig {
	data, err := ioutil.ReadFile("../db_config.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	dbConfig := &DBConfig{}
	err = json.Unmarshal(data, dbConfig)
	if err != nil {
		log.Fatalf("GetDBConfig error:%v", err)
	}

	return dbConfig
}

type DBConfig struct {
	Mysql struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		DBName   string `json:"DBName"`
	} `json:"mysql"`
	Postgresql struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		DBName   string `json:"DBName"`
	} `json:"postgresql"`
	Sqlite3 struct {
		Path string `json:"Path"`
	} `json:"sqlite3"`
}
