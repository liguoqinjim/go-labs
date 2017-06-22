package util

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadXml(result interface{}, name string) {
	f, err := os.Open("gamedata/" + name)

	defer f.Close()

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(data, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println("read xml data success")
}
