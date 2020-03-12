package main

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//lab001()
	//lab002()
	lab003()
}

func lab001() {
	config, _ := toml.Load(`
[postgres]
user = "pelletier"
password = "mypassword"`)
	// retrieve data directly
	user := config.Get("postgres.user").(string)
	log.Println("user=", user)

	// or using an intermediate object
	postgresConfig := config.Get("postgres").(*toml.Tree)
	password := postgresConfig.Get("password").(string)
	log.Println("password=", password)
}

func lab002() {
	type Postgres struct {
		User     string
		Password string
	}
	type Config struct {
		Postgres Postgres
	}

	doc := []byte(`
[Postgres]
User = "pelletier"
Password = "mypassword"`)

	config := Config{}
	toml.Unmarshal(doc, &config)
	log.Println("user=", config.Postgres.User)
}

func lab003() {
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	type Config struct {
		Port  int
		Title string
		Owner struct {
			Name string
			Dob  time.Time
		}
		Database struct {
			Server        string
			Ports         []int
			ConnectionMax int `toml:"connection_max"`
			Enabled       bool
		}
		Servers map[string]struct {
			IP string //注意:这里Ip IP都是可以解析的
			DC string
		}
		Clients struct {
			Data [][]interface{}
		}
		Server struct {
			Port int
		}
		Hosts []interface{}
	}

	config := new(Config)
	if err := toml.Unmarshal(data, config); err != nil {
		log.Fatalf("toml.Unmarshal error:%v", err)
	}
	log.Printf("%+v", config)
}
