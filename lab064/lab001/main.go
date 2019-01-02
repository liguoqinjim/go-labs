package main

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"io/ioutil"
	"log"
)

var connectInfo = &ConnectInfo{}

func init() {
	readConf()
}

func main() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%s", connectInfo.Hostname, connectInfo.Port)},
		Username: connectInfo.Username,
		Password: connectInfo.Pwd,
		Database: connectInfo.DB,
	})
	if err != nil {
		log.Fatalf("mgo.DialWithInfo error:%v", err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	if err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"}); err != nil {
		log.Fatalf("c.Insert error:%v", err)
	}

	result := Person{}
	if err := c.Find(bson.M{"name": "Ale"}).One(&result); err != nil {
		log.Fatalf("c.Find error:%v", err)
	}

	log.Println("Phone:", result.Phone)
}

type Person struct {
	Name  string
	Phone string
}

func readConf() {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	if err := json.Unmarshal(data, connectInfo); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}
}

type ConnectInfo struct {
	Username string
	Pwd      string
	Hostname string
	Port     string
	DB       string
}
