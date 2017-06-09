package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name  string
	Phone string
}

type ConnectInfo struct {
	Username string
	Pwd      string
	Hostname string
	Port     string
	DB       string
}

func readConf() *ConnectInfo {
	file, err := os.Open("mongo.json")
	handleError(err)

	data, err := ioutil.ReadAll(file)
	handleError(err)

	var ci ConnectInfo
	err = json.Unmarshal(data, &ci)
	handleError(err)
	return &ci
}

var connectInfo *ConnectInfo

func main() {
	//读取连接配置
	connectInfo = readConf()

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%s", connectInfo.Hostname, connectInfo.Port)},
		Username: connectInfo.Username,
		Password: connectInfo.Pwd,
		Database: connectInfo.DB,
	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
