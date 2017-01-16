package main

import (
	"encoding/json"
	"fmt"
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
	data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig = *getDBConfig(data)
	fmt.Printf("%+v\n", dbConfig)
}
