package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	"proto/out"
)

type TestProto struct {
	out.UnimplementedTestServer
}

func (tp *TestProto) TestFunc(ctx context.Context, req *out.TestRequest) (*out.TestResponse, error) {
	msg := req.GetMessage()
	res := strings.ToUpper(fmt.Sprintf("Server Response to %s", msg))

	return &out.TestResponse{Message: res}, nil
}

func main()  {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	out.RegisterTestServer(s, &TestProto{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}