package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/nyogjtrc/littlejp/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	throwForever(conn)
}

func throw(conn *grpc.ClientConn) {
	c := pb.NewJPServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.ThrowMoney(ctx, &pb.ThrowRequest{Amount: 1})
	if err != nil {
		panic(err)
	}
	fmt.Printf("JP Amount: %.2f | Is Winner: %t\n", reply.Amount, reply.IsWinner)

	if reply.IsWinner {
		fmt.Println("You are Winner!!")
		os.Exit(0)
	}
}

func throwForever(conn *grpc.ClientConn) {
	for {
		throw(conn)
		time.Sleep(100 * time.Millisecond)
	}
}
