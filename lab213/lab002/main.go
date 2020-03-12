package main

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//lab001()

	lab002()
}

func lab001() {
	data, err := ioutil.ReadFile("lab001.toml")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	conf := new(testDoc)
	if err := toml.Unmarshal(data, conf); err != nil {
		log.Fatalf("toml.Unmarshal error:%v", err)
	}
	log.Printf("config:%+v", conf)
}

type testDoc struct {
	Title       string            //`toml:"title"`
	NewTitle    string            `toml:"new_title"`
	BasicLists  testDocBasicLists `toml:"basic_lists"`
	SubDocPtrs  []*testSubDoc     `toml:"subdocptrs"`
	BasicMap    map[string]string `toml:"basic_map"`
	Subdocs     testDocSubs       `toml:"subdoc"`
	Basics      testDocBasics     `toml:"basic"`
	SubDocList  []testSubDoc      `toml:"subdoclist"`
	err         int               `toml:"shouldntBeHere"`
	unexported  int               `toml:"shouldntBeHere"`
	Unexported2 int               `toml:"-"`
}

type testDocBasics struct {
	Uint       uint      `toml:"uint"`
	Bool       bool      `toml:"bool"`
	Float32    float32   `toml:"float"`
	Float64    float64   `toml:"float64"`
	Int        int       `toml:"int"`
	String     *string   `toml:"string"`
	Date       time.Time `toml:"date"`
	unexported int       `toml:"shouldntBeHere"`
}

type testDocBasicLists struct {
	Floats  []*float32  `toml:"floats"`
	Bools   []bool      `toml:"bools"`
	Dates   []time.Time `toml:"dates"`
	Ints    []int       `toml:"ints"`
	UInts   []uint      `toml:"uints"`
	Strings []string    `toml:"strings"`
}

type testDocSubs struct {
	Second *testSubDoc `toml:"second"`
	First  testSubDoc  `toml:"first"`
}

type testSubDoc struct {
	Name       string `toml:"name"`
	unexported int    `toml:"shouldntBeHere"`
}

func lab002() {
	data, err := ioutil.ReadFile("lab002.toml")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	conf := new(config)
	if err := toml.Unmarshal(data, conf); err != nil {
		log.Fatalf("toml.Unmarshal error:%v", err)
	}

	log.Printf("%+v", conf)
}

type config struct {
	Debug      bool
	NewTitle   string `toml:"new_title"`
	SwaggerMax string `toml:"swagger_max"`
	Port       int

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string `toml:"db_name"`
	}

	Web struct {
		Port     int
		PageSize int
	}

	WX struct {
		AppId   string
		SignKey string
		Server  string
	}
}
