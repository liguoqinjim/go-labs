package main

import (
	"encoding/json"
	"log"
	"math/big"
	"os"
)

type Message struct {
	Name string
	Body string
	Time *big.Int
}

func main() {
	//marshal object->json
	m := Message{"Alice", "Hello", big.NewInt(1335856677732781)}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("b = %s\n", b)

	//unmarshal json->object
	b2 := []byte(`{"Name":"Bob","Food":"Pickle","Time":13358566777395482781}`)
	var m2 Message
	err = json.Unmarshal(b2, &m2)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("m2 = %+v\n", m2)
	log.Printf("time=%s", m2.Time.String())
}
