package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

var (
	url = "http://kafka-hp-01:9200/"
)

func main() {
	ctx := context.Background()

	//curl -XGET 'http://127.0.0.1:9200/_nodes/http?pretty=true'
	callback := func(node *elastic.NodesInfoNode) bool {
		log.Printf("%+v", node)

		// You can e.g. check node.Attributes to decide if the node will receive requests
		if _, found := node.Attributes["ignore-in-elastic"]; found {
			return false
		}
		return true // will get requests
	}
	client, err := elastic.NewClient(elastic.SetSnifferCallback(callback), elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("elastic.NewClient error:%v", err)
	}

	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Println("code=", code)
	log.Printf("info=%+v", info)
}
