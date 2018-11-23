package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	pb "github.com/nyogjtrc/littlejp/proto"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func printCmdUsage() {
	fmt.Println("please use right cmd")
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 1 {
		printCmdUsage()
		os.Exit(1)
	}

	handlefunc := cmdRouter(os.Args[1])
	if handlefunc != nil {
		handlefunc(conn)
	}
}

func cmdRouter(cmd string) func(*grpc.ClientConn) {
	switch cmd {
	case "throw":
		return throw
	case "throws":
		return throwForever
	case "status":
		return getStatus
	case "latest10":
		return latest10
	case "top10":
		return top10
	default:
		printCmdUsage()
		return nil
	}
}

func throw(conn *grpc.ClientConn) {
	c := pb.NewJPServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.ThrowMoney(ctx, &pb.ThrowRequest{UserId: uint64(rand.Intn(100)), Amount: 1})
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

func getStatus(conn *grpc.ClientConn) {
	c := pb.NewJPServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.GetStatus(ctx, &pb.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func latest10(conn *grpc.ClientConn) {
	c := pb.NewJPServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.Latest10Winner(ctx, &pb.Empty{})
	if err != nil {
		panic(err)
	}

	for i := range reply.Recoreds {
		fmt.Println(i, reply.Recoreds[i])
	}

}

func top10(conn *grpc.ClientConn) {
	c := pb.NewJPServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.Top10Winner(ctx, &pb.Empty{})
	if err != nil {
		panic(err)
	}

	for i := range reply.Recoreds {
		fmt.Println(i, reply.Recoreds[i])
	}

}
