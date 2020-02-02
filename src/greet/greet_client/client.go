package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"synapse/src/greet/greetpb"
	"time"
)

func main() {
	fmt.Println("Hello, I'm a client")

	tls := true
	opts := grpc.WithInsecure()
	if tls {
		creds, sslErr := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	connection, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect : %v\n", err)
	}
	defer func() {
		_ = connection.Close()
	}()

	client := greetpb.NewGreetServiceClient(connection)
	//fmt.Printf("Created client: %f", client)

	doUnary(client)

	//doServerStreaming(client)

	//doClientStreaming(client)

	//doBiDiStreaming(client)

	//doUnaryWithDeadline(client, 5*time.Second)

	//doUnaryWithDeadline(client, 1*time.Second)
}

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mike",
			LastName:  "Tirdea",
		},
	}

	response, err := client.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Respone from Greet: %v", response.Result)
}

func doServerStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server streaming RPC")

	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mihail",
			LastName:  "Tirdea",
		},
	}

	stream, err := client.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Reached end of stream")
			// reached end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}
		log.Printf("Response from GreetManyTimes : %v", message.GetResult())
	}
}

func doClientStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client streaming RPC")

	names := []string{
		"Mihail",
		"John",
		"Bobi",
		"Igor",
		"Gheorghi",
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, name := range names {
		fmt.Printf("Sending name: %v\n", name)
		_ = stream.Send(&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: name,
			},
		})
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet Response: %v\n", response)
}

func doBiDiStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Bidi Streaming RPC")

	names := []string{
		"Mihail",
		"John",
		"Bobi",
		"Igor",
		"Gheorghi",
	}

	// create stream by invoking the client
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating the stream : %v", err)
	}
	waitChan := make(chan struct{})
	// send messages to the client ( go routine )
	go func() {
		// function to send messages
		for _, name := range names {
			fmt.Printf("Sending message: %v\n", name)
			_ = stream.Send(&greetpb.GreetEveryoneRequest{
				Greeting: &greetpb.Greeting{
					FirstName: name,
				},
			})
			//time.Sleep(1000 * time.Millisecond)
		}
		fmt.Println("Finished sending")
		_ = stream.CloseSend()
	}()

	// receive messages from the client ( go routine )
	go func() {
		// function to receive messages
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			fmt.Printf("Received: %v\n", response.GetResult())
		}
		fmt.Println("Finished receiving")
		close(waitChan)
	}()

	<-waitChan // block until everything is done
}

func doUnaryWithDeadline(client greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do a UnaryWithDeadline RPC")

	request := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mihail",
			LastName:  "Tirdea",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := client.GreenWithDeadline(ctx, request)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok { // gRPC error
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("Unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}

		return
	}

	fmt.Printf("Response from GreetWithDeadline: %v\n", response.GetResult())
}
