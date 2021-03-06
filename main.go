package main

import (
	"context"
	"fmt"
	"log"
	"time"

	helloserice "github.com/fudute/proto_idl/helloservice"
	"google.golang.org/grpc"
)

func main() {
	ShortConn()
}

func Default() {
	cc, err := grpc.Dial("hello:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	cli := helloserice.NewHelloServiceClient(cc)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Hello(ctx, &helloserice.HelloRequest{Name: "fudute"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("resp: %v\n", resp)
		time.Sleep(time.Second)

		cancel()
	}
}

func ShortConn() {
	for {
		cc, err := grpc.Dial("hello:8080", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatal(err)
		}

		cli := helloserice.NewHelloServiceClient(cc)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Hello(ctx, &helloserice.HelloRequest{Name: "fudute"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("resp: %v\n", resp)
		time.Sleep(time.Second)
		cancel()
	}
}
