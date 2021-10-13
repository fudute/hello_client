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
	cc, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
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
