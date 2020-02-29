package main

import (
	"context"
	"log"
	"net/http"
	"os"

	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
)

var events = neffos.Namespaces{
	"v1": neffos.Events{
		"echo": onEcho,
	},
}

func onEcho(c *neffos.NSConn, msg neffos.Message) error {
	body := string(msg.Body)
	log.Println(body)

	if !c.Conn.IsClient() {
		newBody := append([]byte("echo back: "), msg.Body...)
		return neffos.Reply(newBody)
	}

	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("expected program to start with 'server' or 'client' argument")
	}
	side := args[0]

	switch side {
	case "server":
		runServer()
	case "client":
		runClient()
	default:
		log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
	}
}

func runServer() {
	//websocketServer := neffos.New(gorilla.DefaultUpgrader, events)

	websocketServer := neffos.New(
		gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}),
		events)

	router := http.NewServeMux()
	router.Handle("/echo", websocketServer)

	log.Println("Serving websockets on localhost:8080/echo")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func runClient() {
	ctx := context.Background()
	client, err := neffos.Dial(ctx, gorilla.DefaultDialer, "ws://localhost:8080/echo", events)
	if err != nil {
		panic(err)
	}

	c, err := client.Connect(ctx, "v1")
	if err != nil {
		panic(err)
	}

	c.Emit("echo", []byte("Greetings!"))

	// a channel that blocks until client is terminated,
	// i.e by CTRL/CMD +C.
	<-client.NotifyClose
}
