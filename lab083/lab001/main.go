package main

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"gopkg.in/olivere/elastic.v5"
	"io/ioutil"
	"log"
)

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

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

	// 创建一个连接elasticsearch的client，默认是连接127.0.0.1:9200，也可以设置连接url
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Fatalf("NewClient error:%v", err)
	}

	// Ping Elasticsearch，可以得到version number等数据
	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		log.Fatalf("client.Ping error:%v", err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// 得到es version
	esversion, err := client.ElasticsearchVersion(url)
	if err != nil {
		log.Fatalf("ElasticsearchVersion error:%v", err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)

	// 检查index是否存在
	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists error:%v", err)
	}
	if !exists {
		// 创建一个新的index
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			log.Fatalf("CreateIndex error:%v", err)
		}
		log.Printf("CreateIndex result,Acknowledged[%t]ShardsAckonwledged[%t]", createIndex.Acknowledged, createIndex.ShardsAcknowledged)
	} else {
		log.Println("index existed")
	}

	// Index a tweet (using JSON serialization)
	// 就是存入一个document(相当于mysql里面的insert一行数据)
	tweet1 := Tweet{User: "kimi", Message: "This is the first message", Retweets: 0}
	put1, err := client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet1).
		Do(ctx)
	if err != nil {
		log.Fatalf("Index error:%v", err)
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Index a second tweet (by string)
	// 存入第二个document
	tweet2 := `{"user" : "kimi", "message" : "This is the second message"}`
	put2, err := client.Index().
		Index("twitter").
		Type("tweet").
		Id("2").
		BodyString(tweet2).
		Do(ctx)
	if err != nil {
		log.Fatalf("Index error:%v", err)
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Get tweet with specified ID
	// 通过id得到document
	get1, err := client.Get().
		Index("twitter").
		Type("tweet").
		Id("1").
		Do(ctx)
	if err != nil {
		log.Fatalf("Get error:%v", err)
	}
	if get1.Found {
		log.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index("twitter").Do(ctx)
	if err != nil {
		log.Fatalf("Flush error:%v", err)
	}

	// Search with a term query
	// 查询
	termQuery := elastic.NewTermQuery("user", "kimi")
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		log.Fatalf("Search error:%v", err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	// searchResult里面会有一些查询结果的数据
	log.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	// 查询结果
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			log.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	log.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		log.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			log.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		log.Print("Found no tweets\n")
	}

	// Update a tweet by the update API of Elasticsearch.
	// We just increment the number of retweets.
	// 更新
	update, err := client.Update().Index("twitter").Type("tweet").Id("1").
		Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do(ctx)
	if err != nil {
		log.Fatalf("Update error:%v", err)
	}
	log.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)

	// Delete an index.
	// 删除index
	deleteIndex, err := client.DeleteIndex("twitter").Do(ctx)
	if err != nil {
		log.Fatalf("client.Delete error:%v", err)
	} else {
		if !deleteIndex.Acknowledged {
			// Not acknowledged
		}

		log.Printf("index twitter deleted")
	}
}
