package main

import (
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"log"
)

func main() {
	var conns []bolt.Conn

	driver := bolt.NewDriver()
	for i := 0; i < 5; i++ {
		conn, err := driver.OpenNeo("bolt://neo4j:123456@hp-113:7687")
		if err != nil {
			log.Fatalf("driver.OpenNeo error:%v", err)
		}

		conns = append(conns, conn)
	}

	for i := 0; i < 5; i++ {
		log.Println("conn", i)
		conn := conns[i]

		data, _, _, err := conn.QueryNeoAll("MATCH p=()-->() RETURN p LIMIT 25", nil)
		if err != nil {
			log.Fatalf("conn.QueryNeoAll error:%v", err)
		}
		for _, row := range data {
			log.Printf("NODE: %#v\n", row[0].(graph.Path)) // Prints all nodes
		}
	}

	for _, v := range conns {
		v.Close()
	}
}
