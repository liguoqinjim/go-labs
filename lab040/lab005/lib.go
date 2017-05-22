package lib

import (
	"errors"
	"fmt"
	pb "lab040/lab005/message"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	rpcPort     = ":50050"
	grpcPort    = ":50051"
	grpcAddress = "localhost:50051"
)

type GRPCCalServer struct{}

func (s *GRPCCalServer) Mult(ctx context.Context, in *pb.CalRequest) (*pb.MultResult, error) {
	result := in.A * in.B
	return &pb.MultResult{Result: result}, nil
}
func (s *GRPCCalServer) Div(ctx context.Context, in *pb.CalRequest) (*pb.DivResult, error) {
	if in == nil {
		return nil, errors.New("nil param")
	}
	if in.B == 0 {
		return nil, errors.New("b cannot be 0")
	}

	divResult := new(pb.DivResult)
	divResult.Result = in.A / in.B
	divResult.Mod = in.A % in.B
	return divResult, nil
}

func NewGRPCServer() {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &GRPCCalServer{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve:%v\n", err)
	}
}

func NewGRPCClient() {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect:%v\n", err)
	}
	defer conn.Close()

	// client := pb.NewCalculatorClient(conn)

	// multRequest := &pb.CalRequest{A: 2, B: 3}

}
