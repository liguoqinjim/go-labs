package main

import (
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
	"net/http"
)

type clientPage struct {
	Title string
	Host  string
}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".html"))

	ws := websocket.New(gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), websocket.Events{
		websocket.OnNativeMessage: func(conn *neffos.NSConn, message neffos.Message) error {
			log.Printf("Server got: %s from [%s]", message.Body, conn.Conn.ID())

			conn.Conn.Server().Broadcast(conn, message)
			return nil
		},
	})

	ws.OnConnect = func(c *neffos.Conn) error {
		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *neffos.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
	}

	app.HandleDir("/js", "./static/js")
	app.Get("/my_endpoint", websocket.Handler(ws))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("client.html", clientPage{"Client Page", "localhost:8080"})
	})

	app.Run(iris.Addr(":8080"))
}
