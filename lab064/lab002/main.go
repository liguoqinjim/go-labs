package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type ConnectInfo struct {
	Username string
	Pwd      string
	Hostname string
	Port     string
	DB       string
}

var connectInfo *ConnectInfo

func readConf() *ConnectInfo {
	file, err := os.Open("mongo.json")
	handleError(err)

	data, err := ioutil.ReadAll(file)
	handleError(err)

	ci := new(ConnectInfo)
	err = json.Unmarshal(data, ci)
	handleError(err)

	return ci
}

type Game struct {
	ID           bson.ObjectId `bson:"_id,omitempty"` //_id这样可以收到mongodb的id，omitempty可以在insert时候不插入这个值，而是由mongodb自动生成
	Winner       string        `bson:"winner"`
	OfficialGame bool          `bson:"official_game"`
	Location     string        `bson:"location"`
	StartTime    time.Time     `bson:"start"`
	EndTime      time.Time     `bson:"end"`
	Players      []Player      `bson:"players"`
}

type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
	return Player{
		Name:   name,
		Decks:  [2]string{firstDeck, secondDeck},
		Points: points,
		Place:  place,
	}
}

const colName = "lab002"
const dbName = "test"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	connectInfo = readConf()

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%s", connectInfo.Hostname, connectInfo.Port)},
		Username: connectInfo.Username,
		Password: connectInfo.Pwd,
		Database: connectInfo.DB,
	})
	handleError(err)
	defer session.Close()

	gameA := Game{
		Winner:       "Dave",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Date(2015, time.February, 12, 05, 54, 0, 0, time.UTC),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}

	gameB := Game{
		Winner:       "Javier",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Date(2015, time.February, 12, 05, 54, 0, 0, time.UTC),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//drop collection
	err = session.DB(dbName).C(colName).DropCollection()
	if err != nil {
		log.Println("dropCollection error", err)
	}

	//insert document
	coll := session.DB(dbName).C(colName)
	if err = coll.Insert(gameA); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert success")
	}
	if err = coll.Insert(gameB); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert success")
	}

	//read document
	official_game := true
	gamesWonCount, err := coll.Find(bson.M{"official_game": official_game}).Count()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("gamesWonCount=", gamesWonCount)
	}

	//read document
	var gamesWon *Game
	playerName := "Dave"
	err = coll.Find(bson.M{"winner": playerName}).One(&gamesWon)
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Printf("gamesWon=%+v\n", gamesWon)
		fmt.Println("gamesWon.Winner=", gamesWon.Winner)
	}

	//updating a document
	newWinner := "Seth"
	update := bson.M{"$set": bson.M{"winner": newWinner}}
	if err := coll.UpdateId(gamesWon.ID, update); err != nil {
		log.Fatal(err)
	}
	//check update result
	err = coll.FindId(gamesWon.ID).One(&gamesWon)
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Println("gamesWon=%+v\n", gamesWon)
		fmt.Println("gamesWon.Winner=", gamesWon.Winner)
	}

	//deleting a document
	info, err := coll.RemoveAll(bson.M{"winner": gamesWon.Winner})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("remove info.Removed=", info.Removed) //info.Removed就是删除了几个
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
