package main

import (
	"context"
	"fmt"
	"grpc-simple/pkg/api"
	"log"

	"google.golang.org/grpc"
)

const address = ":7000"

//Simple client
func main() {
	var x, y int32
	fmt.Print("Enter x: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter y: ")
	_, err = fmt.Scan(&y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("x: %d, y: %d\n", x, y)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewMathClient(conn)
	fmt.Println("Client is ready...")

	fmt.Println("Calling rpc Sum...")
	sumResp, err := client.Sum(context.Background(), &api.SumReq{X: x, Y: y})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum result: ", sumResp.Result)

	fmt.Println("Calling rpc Sub...")
	subResp, err := client.Sub(context.Background(), &api.SubReq{X: x, Y: y})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sub result: ", subResp.Result)
}
