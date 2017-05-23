package lib_test

import (
	"lab040/lab005/lib"
	pb "lab040/lab005/message"
	"log"
	"net"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func BenchmarkGRPCMain(b *testing.B) {
	b.StopTimer()
	log.Println(time.Now())
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

	time.Sleep(time.Second * 2)
	log.Println(time.Now())
	lib.NewGRPCClient()

	b.StartTimer()
	b.Run("mult", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			multResult, err := lib.GRPCCalClient.Mult(context.Background(), &pb.CalRequest{A: 2, B: 3})
			if err != nil {
				log.Fatal("mult error ", err)
			}
			if multResult.Result != 6 {
				log.Fatal("cal mult error ", multResult.Result)
			}
		}
	})

	b.Run("div", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			divResult, err := lib.GRPCCalClient.Div(context.Background(), &pb.CalRequest{A: 7, B: 3})
			if err != nil {
				log.Fatal("div error", err)
			}
			if divResult.Result != 2 || divResult.Mod != 1 {
				log.Fatal("cal div error ", divResult)
			}
		}
	})
}

func BenchmarkRPCMain(b *testing.B) {
	b.StopTimer()
	s := lib.NewRPCServer()
	listener, err := net.Listen("tcp", lib.RPCPort)
	if err != nil {
		log.Fatal("listener error", err)
	}
	go func() {
		s.Accept(listener)
	}()

	time.Sleep(time.Second * 2)
	lib.NewRPCClient()

	b.StartTimer()
	//计算
	b.Run("rpcMult", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var multResult int
			err := lib.RPCCalClient.Call("MyMethod.Mult", &lib.RArgs{A: 7, B: 3}, &multResult)
			if err != nil {
				log.Fatal("mult error", err)
			}
			if multResult != 21 {
				log.Fatal("multResult error", multResult)
			}
		}
	})

	b.Run("rpcDiv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reply := new(lib.RDivResult)
			err := lib.RPCCalClient.Call("MyMethod.Div", &lib.RArgs{A: 7, B: 3}, reply)
			if err != nil {
				log.Fatal("div error", err)
			}
			if reply.Result != 2 || reply.Mod != 1 {
				log.Fatal("divResult error", reply)
			}
		}
	})
}
