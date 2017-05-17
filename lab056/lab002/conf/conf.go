package conf

import (
	"io/ioutil"
	"encoding/json"
)

var ConnConf struct{
	Addr string
	Password string
	DB string
}

func ReadConf(){
	b,err := ioutil.ReadFile("conf.json")
	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(b,&ConnConf)
	if err != nil{
		panic(err)
	}
}
