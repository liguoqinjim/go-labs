package conf

import (
	"encoding/json"
	"io/ioutil"
)

var ConnConf struct {
	Addr     string
	Password string
	DB       int
}

func ReadConf() {
	b, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &ConnConf)
	if err != nil {
		panic(err)
	}
}
