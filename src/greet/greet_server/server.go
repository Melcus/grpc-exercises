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
	"net"
	"strconv"
	"synapse/src/greet/greetpb"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	response := &greetpb.GreetResponse{
		Result: result,
	}

	return response, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		response := &greetpb.GreetManyTimesResponse{

			Result: result,
		}
		_ = stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request \n")
	result := ""
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("We have finished reading the client stream")
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := request.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with a streaming request \n")

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
			return err
		}

		firstName := request.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! \n"

		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func (*server) GreenWithDeadline(ctx context.Context, request *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreenWithDeadline function was invoked with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client cancelled the request
			fmt.Println("The client cancelled the request")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	result := "Hello " + firstName
	response := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return response, nil
}

func main() {
	fmt.Println("Starting the server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tls := true
	var opts []grpc.ServerOption
	if tls {
		creds, sslErr := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if sslErr != nil {
			log.Fatalf("Failed loading certificates, %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
