package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"synapse/src/calculator/calculatorpb"
	"time"
)

func main() {
	fmt.Println("Calculator client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v\n", err)
	}
	defer func() {
		_ = connection.Close()
	}()

	client := calculatorpb.NewCalculatorServiceClient(connection)
	//fmt.Printf("Created client: %f", client)

	//doUnary(client)

	//doServerStreaming(client)

	//doClientStreaming(client)

	doBiDiStreaming(client)
}

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecompositionServer Streaming RPC")
	request := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 1235943,
	}
	stream, err := client.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("End of transmission")
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(response.GetPrimeFactor())
	}
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC")
	request := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Respone from Greet: %v", response.SumResult)
}

func doClientStreaming(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client streaming RPC")

	numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v", err)
	}

	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		_ = stream.Send(&calculatorpb.AverageRequest{
			Number: number,
		})
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v", err)
	}

	fmt.Printf("Average Response: %v\n", response)
}

func doBiDiStreaming(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a FindMaximum BiDi streaming RPC")

	stream, err := client.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	waitChan := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{1, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("Sending number %v\n", number)
			_ = stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		fmt.Println("Finished sending")
		_ = stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Server stopped streaming")
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v", err)
			}

			maximum := response.GetMaximum()
			fmt.Printf("Received a new maximum of ...: %v\n", maximum)
		}
		fmt.Println("Finished receiving")
		close(waitChan)
	}()

	<-waitChan
}
