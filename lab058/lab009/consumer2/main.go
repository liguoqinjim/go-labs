package main

import (
	"context"
	"github.com/nsqio/go-nsq"
	elastic "gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	INDEX_NAME     = "server900002"
	NSQ_TOPIC_NAME = "server900001"
)

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
				},
				"logType1":{
					"type":"keyword"
				},
				"logType2":{
					"type":"keyword"
				}
			}
		}
	}
}`

type ActionLog struct {
	PlayerId string `json:"playerId"`
	Cmd      string `json:"cmd"`
	LogType1 string `json:"logType1"`
	LogType2 string `json:"logType2"`
	LogTime  string `json:"LogTime"`
	Content  string `json:"content"`
}

var (
	ELKClient  *elastic.Client
	ELKContext context.Context
)

func init() {
	ConnectToElastic()
}

func main() {
	data, err := ioutil.ReadFile("ip.conf")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}

	addr := string(data)

	//创建连接
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(NSQ_TOPIC_NAME, "ch", config)
	if err != nil {
		log.Fatalf("NewConsumer error:%v", err)
	}
	q.AddHandler(nsq.HandlerFunc(MsgHandler))

	//连接
	err = q.ConnectToNSQLookupd(addr)
	if err != nil {
		log.Fatalf("ConnnectToNSQLookupd error:%v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Println("end")
}

func ConnectToElastic() {
	data, err := ioutil.ReadFile("elk.conf")
	if err != nil {
		log.Fatalf("ReadFile elk error:%v", err)
	}
	url := string(data)

	ELKContext = context.Background()

	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Fatalf("NewClient error:%v", err)
	}

	info, code, err := client.Ping(url).Do(ELKContext)
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(url)
	if err != nil {
		log.Fatalf("ElasticsearchVersion error:%v", err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(INDEX_NAME).Do(ELKContext)
	if err != nil {
		log.Fatalf("IndexExists error:%v", err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(INDEX_NAME).BodyString(mapping).Do(ELKContext)
		if err != nil {
			log.Fatalf("CreateIndex error:%v", err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	ELKClient = client
}

func AddDocument(actionLog *ActionLog) {
	_, err := ELKClient.Index().
		Index(INDEX_NAME).
		Type("ActionLog").BodyJson(actionLog).Do(ELKContext)
	if err != nil {
		log.Fatalf("Index error:%v", err)
	} else {
		log.Println("添加document")
	}
}

func MsgHandler(message *nsq.Message) error {
	logMessage := string(message.Body)
	logMessages := strings.Split(logMessage, "|")
	log.Println("消息处理:", string(message.ID[:]), string(message.Body))

	if len(logMessages) == 6 || len(logMessages) == 5 {
		actionLog := new(ActionLog)
		actionLog.PlayerId = logMessages[0]
		actionLog.Cmd = logMessages[1]
		actionLog.LogType1 = logMessages[2]
		actionLog.LogType2 = logMessages[3]
		actionLog.LogTime = logMessages[4]

		//判断log类型
		AddDocument(actionLog)
	}

	return nil
}
