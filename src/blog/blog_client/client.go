package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"synapse/src/blog/blogpb"
)

func main() {
	fmt.Println("Blog client")

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

	client := blogpb.NewArticleServiceClient(connection)

	fmt.Println("Creating the article")
	article := &blogpb.Article{
		AuthorId: "Mike",
		Title:    "First article",
		Content:  "Content of the first article :)",
	}

	createArticleResponse, err := client.CreateArticle(context.Background(), &blogpb.CreateArticleRequest{Article: article})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Article has been created: %v\n", createArticleResponse)
}
