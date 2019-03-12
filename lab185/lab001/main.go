package main

import (
	"encoding/json"
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	dbConfig := readConf()
	log.Println("dbConfig=", dbConfig.Mysql)

	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%d", dbConfig.Mysql.Host, dbConfig.Mysql.Port)
	cfg.User = dbConfig.Mysql.User
	cfg.Password = fmt.Sprintf("'%s'", dbConfig.Mysql.Password)

	cfg.Dump.TableDB = "hp_wifi"
	cfg.Dump.Tables = []string{"t_heat_map_hour"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatalf("canal.NewCanal error:%v", err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	if err := c.Run(); err != nil {
		log.Fatalf("c.Run error:%v", err)
	}

	time.Sleep(time.Hour)
}

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
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
