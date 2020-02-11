package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"os/signal"
	"synapse/src/blog/blogpb"
	"time"
)

type server struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog Service started")

	// Init MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://synapse:secret@localhost:27030"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB server: %v", err)
	}else {
		fmt.Println("MongoDB ping OK !")
	}

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
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	_ = listener.Close()
	fmt.Println("End of program")
}
