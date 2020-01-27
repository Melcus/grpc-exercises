package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"synapse/src/calculator/calculatorpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC : %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber

	result := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return result, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			_ = stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	fmt.Printf("Average function was invoked with a streaming request \n")
	result := int64(0)
	count := 0
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("We have finished reading the client stream")
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Average: float64(result) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		result += request.GetNumber()
		count++
	}
}

func main() {
	fmt.Println("Calculator Server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
