package lib

import (
	"errors"
	pb "lab040/lab005/message"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	RPCPort     = ":50050"
	GRPCPort    = ":50051"
	GRPCAddress = "localhost:50051"
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

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &GRPCCalServer{})
	reflection.Register(s)

	return s
}

var GRPCCalClient pb.CalculatorClient

func NewGRPCClient() {
	conn, err := grpc.Dial(GRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect:%v\n", err)
	}
	defer conn.Close()
	GRPCCalClient = pb.NewCalculatorClient(conn)
}
