package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	client *mongo.Client
)

var (
	host     string
	port     string
	user     string
	password string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pflag.StringVarP(&host, "host", "h", "127.0.0.1", "db host")
	pflag.StringVarP(&port, "port", "p", "27017", "mongo db port")
	pflag.StringVarP(&user, "user", "u", "root", "db user")
	pflag.StringVarP(&password, "password", "P", "", "db port")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	open()
	findOne()
}

func open() {
	//mongodb://root@localhost:27017/?authSource=admin
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin", user, password, host, port)
	log.Printf("uri=%s", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	var err error
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("mongo.Connect error:%v", err)
	} else {
		log.Println("mongo.Connect success")
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		log.Fatalf("disconnect error:%v", err)
	//	}
	//}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("client.Ping error:%v", err)
	} else {
		log.Println("Successfully connected and pinged.")
	}
}

func findOne() {
	db := client.Database("recommender")

	cols, err := db.ListCollections(context.Background(), bson.D{})
	if err != nil {
		log.Fatalf("db.ListCollections error:%v", err)
	}
	for cols.Next(context.Background()) {
		log.Println(cols.Current.String())
	}

	//DecodeBytes
	col := db.Collection("Tag")
	result, err := col.FindOne(context.Background(), bson.D{}).DecodeBytes()
	if err != nil {
		log.Fatalf("col.FindOne error:%v", err)
	}
	log.Println(result.String())

	//Decode
	var tag Tag
	if err := col.FindOne(context.Background(), bson.D{}).Decode(&tag); err != nil {
		log.Fatalf("col.FindOne Decode error:%v", err)
	}
	log.Printf("tag=%+v", tag)
}

type Tag struct {
	ID        primitive.ObjectID `bson:"_id"`
	UID       int                `bson:"uid"`
	Mid       int                `bson:"mid"`
	Tag       string             `bson:"tag"`
	Timestamp int                `bson:"timestamp"`
}
