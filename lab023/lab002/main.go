package main

import (
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	dialer := &websocket.Dialer{}

	conn, _, err := dialer.Dial("wss://io.sosobtc.com/socket.io/?EIO=3&transport=websocket",
		map[string][]string{"Origin": []string{"https://k.sosobtc.com"},
			"Cookie":                []string{"OID=aEj%252BuelTgv0RAuNv%252FFJPBfslpUnvZ26EiWsVM7TiIgNZ%252FaJQtLXiwjyAsqbnRKA%252BpQ7UkYv1rrO92kq8%252BZ4ifZQex9e7Sbgj7BVy3DtSflfIJd4koi1JTx61ElPwSY8x%7C8dad2860013668cf3e1c4aa6c4e19154; _ga=GA1.2.241055403.1496717869; _gid=GA1.2.1622829335.1496717869; _gat=1; theme=dark"},
			"User-Agent":            []string{"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"},
			"Sec-WebSocket-Version": []string{"13"}})
	if err != nil {
		log.Fatal("error= ", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		defer conn.Close()
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	time.Sleep(time.Millisecond * 500)
	sendMessage := `42["channel.ticker.subscribe",["btc:btc38","btc:poloniex","btc:bittrex"]]`
	err = conn.WriteMessage(websocket.TextMessage, []byte(sendMessage))
	if err != nil {
		log.Println("write:", err)
		return
	}

	sendMessage2 := `4212["market.subscribe","sc:yunbi"]`
	err = conn.WriteMessage(websocket.TextMessage, []byte(sendMessage2))
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			conn.Close()
			return
		}
	}
}
