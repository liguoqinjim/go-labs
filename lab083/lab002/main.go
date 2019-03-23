package main

import (
	"context"

	"fmt"
	elastic "github.com/olivere/elastic"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ActionLog struct {
	PlayerId int    `json:"playerId"`
	Cmd      int    `json:"cmd"`
	Content  string `json:"content"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"actionLog":{
			"properties":{
				"playerId":{
					"type":"keyword"
				},
				"cmd":{
					"type":"keyword"
				}
			}
		}
	}
}`

const (
	INDEX_NAME = "server900001"
	TYPE_NAME  = "actionLog"
)

func main() {
	//读取ip
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("ReadFile error:%v", err)
	}
	url := string(data)
	log.Printf("url=%s\n", url)

	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Fatalf("NewClient error:%v", err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(url)
	if err != nil {
		log.Fatalf("ElasticsearchVersion error:%v", err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(INDEX_NAME).Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists error:%v", err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(INDEX_NAME).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Fatalf("CreateIndex error:%v", err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Index a tweet (using JSON serialization)
	actionLog1 := ActionLog{PlayerId: 68, Cmd: 2088, Content: "行为日志"}
	put1, err := client.Index().
		Index(INDEX_NAME).
		Type(TYPE_NAME).
		Id("1").
		BodyJson(actionLog1).
		Do(ctx)
	if err != nil {
		log.Fatalf("Index error:%v", err)
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Index a second tweet (by string)
	actionLog2 := `{"playerId":67,"cmd":2008,"content":"行为日志2"}`
	put2, err := client.Index().
		Index(INDEX_NAME).
		Type(TYPE_NAME).
		Id("2").
		BodyString(actionLog2).
		Do(ctx)
	if err != nil {
		log.Fatalf("Index error:%v", err)
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Get tweet with specified ID
	get1, err := client.Get().
		Index(INDEX_NAME).
		Type(TYPE_NAME).
		Id("1").
		Do(ctx)
	if err != nil {
		log.Fatalf("Get error:%v", err)
	}
	if get1.Found {
		log.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index(INDEX_NAME).Do(ctx)
	if err != nil {
		log.Fatalf("Flush error:%v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		playerIds := []int{4, 11, 68, 70}
		cmds := []int{2001, 2003, 2005, 2401, 2088, 2009, 2006, 2409, 2004, 2007, 2002}

		i := 0
		for {
			playerId := playerIds[rand.Intn(len(playerIds))]
			cmd := cmds[rand.Intn(len(cmds))]

			actionLog := ActionLog{PlayerId: playerId, Cmd: cmd, Content: fmt.Sprintf("行为日志%d", i)}

			_, err := client.Index().
				Index(INDEX_NAME).
				Type(TYPE_NAME).
				BodyJson(actionLog).
				Do(ctx)
			if err != nil {
				log.Fatalf("Index error:%v", err)
			} else {
				log.Printf("insert success")
			}

			i++
			time.Sleep(time.Second * 3)
		}
	}()

	<-sigs
	_, err = client.Flush().Index(INDEX_NAME).Do(ctx)
	if err != nil {
		log.Fatalf("Flush error:%v", err)
	}

	log.Println("end")
}
