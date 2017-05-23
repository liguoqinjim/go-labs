package lib_test

import (
	//"golang.org/x/net/context"
	"lab040/lab005"
	//pb "lab040/lab005/message"
	"log"
	"net"
	"testing"
	"time"
)

func BenchmarkGRPC(b *testing.B) {
	b.StopTimer()

	//打开服务器
	s := lib.NewGRPCServer()
	listener, err := net.Listen("tcp", lib.GRPCPort)
	if err != nil {
		log.Fatal("listener error", err)
	}

	go func() {
		err = s.Serve(listener)
		if err != nil {
			log.Fatal("serve error", err)
		}
	}()

	time.Sleep(time.Second)
	lib.NewGRPCClient()

	b.StartTimer()

	//开始计算
	for i := 0; i < b.N; i++ {
		//multResult, err := lib.GRPCCalClient.Mult(context.Background(), &pb.CalRequest{A: 2, B: 3})
		//if err != nil {
		//	log.Fatal("mult error", err)
		//}
		//if multResult.Result != 6 {
		//	log.Fatal("cal mult error", multResult.Result)
		//}
		a := 1 + 1
		if a != 2 {
			log.Fatal("errr")
		}
	}
}

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 1 + 1
		if a != 2 {
			log.Fatal("errr")
		}
	}
}
