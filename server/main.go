package main

import (
	"fmt"
	"net"

	"github.com/nyogjtrc/littlejp"
	pb "github.com/nyogjtrc/littlejp/proto"
	"google.golang.org/grpc"
)

const LISTEN_ADDRESS = ":50051"

func main() {
	fmt.Println("listen:", LISTEN_ADDRESS)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterJPServiceServer(s, littlejp.NewServer())

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
