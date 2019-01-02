package main

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"io/ioutil"
	"log"
)

func main() {
	example()
}

func example() {
	conf := readConf()
	if conf == nil {
		log.Fatalf("conf is nil")
	}

	//打开连接
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%s", conf.Host, conf.Port)},
		Username: conf.Username,
		Password: conf.Password,
		Database: conf.Db,
	})
	if err != nil {
		log.Fatalf("mgo.DialWithInfo error:%v", err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("go-labs-test").C("people")
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

//读取配置文件
func readConf() *Conf {
	data, err := ioutil.ReadFile("../conf.json")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	var conf = &Conf{}
	if err := json.Unmarshal(data, conf); err != nil {
		log.Fatalf("json.Unmarshal error:%v", err)
	}
	return conf
}

type Conf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Db       string `json:"db"`
}
