package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/PhilShaughnes/grpcalc/pb"
)

func main() {
	serverAddr := flag.String(
		"server",
		"localhost:8080",
		"The server address formatted as host:port",
	)
	flag.Parse()

	creds := insecure.NewCredentials()
	// creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{10, 5, 3, 1},
	})
	// res, err := client.Divide(ctx, &pb.CalculationRequest{
	// 	A: 10,
	// 	B: 0,
	// })
	if err != nil {
		log.Fatalln("error sending request:", err)
	}

	fmt.Println("ans:", res.Result)
	fmt.Println("hello world")
}
