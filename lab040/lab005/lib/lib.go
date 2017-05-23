package lib

import (
	"errors"
	pb "lab040/lab005/message"
	"log"

	"net/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	RPCPort     = ":50050"
	RPCAddress  = "127.0.0.1:50050"
	GRPCPort    = ":50051"
	GRPCAddress = "127.0.0.1:50051"
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
	// defer conn.Close()
	GRPCCalClient = pb.NewCalculatorClient(conn)
}

//rpc
type MyMethod int

type RArgs struct {
	A int
	B int
}
type RDivResult struct {
	Result int
	Mod    int
}

func (m *MyMethod) Mult(args *RArgs, reply *int) error {
	if args == nil || reply == nil {
		return errors.New("nil param")
	}

	*reply = args.A * args.B
	return nil
}

func (m *MyMethod) Div(args *RArgs, reply *RDivResult) error {
	if args == nil || reply == nil {
		return errors.New("nil param")
	}

	if args.B == 0 {
		return errors.New("B cannot be 0")
	}

	reply.Result = args.A / args.B
	reply.Mod = args.A % args.B

	return nil
}

func NewRPCServer() *rpc.Server {
	mm := new(MyMethod)
	server := rpc.NewServer()
	server.Register(mm)

	return server
}

var RPCCalClient *rpc.Client

func NewRPCClient() {
	r, err := rpc.Dial("tcp", RPCAddress)
	if err != nil {
		log.Fatal("dial error", err)
	}
	RPCCalClient = r
}
