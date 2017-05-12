package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"os"
)

func main() {
	testjson := ""
	file, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	testjson = string(chunks)

	//const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//value1 := gjson.Get(json, "name.last")
	//println(value1.String())
	//value2 := gjson.Get(json, "age")
	//println(value2.Int())

	//测试1
	//value1 := gjson.Get(testjson, "Back.Params.ArmyGroup1Init.Armys")
	//for _, v := range value1.Array() {
	//	value2 := gjson.Get(v.String(), "PosX")
	//	fmt.Println(value2)
	//}

	matchPath := fmt.Sprintf("Back.Params.BattleFrameDatas.#[Frame==\"%d\"]#", 10)
	fmt.Println(matchPath)
	value1 := gjson.Get(testjson, matchPath)
	fmt.Println(value1)
	fmt.Println(len(value1.Array()))
}
