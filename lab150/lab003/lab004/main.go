package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kataras/iris"
	"log"
)

func main() {
	app := iris.New()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(err error) {
		fmt.Println("meet error:", err)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	app.HandleMany("GET POST", "/socket.io/{any:path}", iris.FromStd(server))
	app.HandleDir("/", "./asset")
	app.Run(iris.Addr(":8000"),
		iris.WithoutPathCorrection,
		iris.WithoutServerError(iris.ErrServerClosed))
}
