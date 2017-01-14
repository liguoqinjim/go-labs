package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
)

func main() {
	//string->object
	var testJson = []byte(`[
		{"Name": "Alice", "Age": 13},
		{"Name": "Bob",    "Age": 15}
	]`)

	json1, err := simplejson.NewJson(testJson)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json2, err := json1.Array()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < len(json2); i++ {
		name, err := json1.GetIndex(i).Get("Name").String()
		if err != nil {
			fmt.Println(1, err)
			break
		}
		age, err := json1.GetIndex(i).Get("Age").Int()
		if err != nil {
			fmt.Println(2, err)
			break
		}
		fmt.Printf("name=%s,age=%d\n", name, age)
	}

	//object->string
	fmt.Println()
	js := simplejson.New()
	js.Set("Name", "Bruce")
	js.Set("Age", 20)
	b, err := js.Encode()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("b = %s\n", b)
}
