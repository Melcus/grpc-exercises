package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
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
	articleId := createArticleResponse.GetArticle().GetId()

	// read article
	fmt.Println("Retrieving the article")

	_, intentionalErr := client.ReadArticle(context.Background(), &blogpb.ReadArticleRequest{ArticleId: "aaasas"})
	if intentionalErr != nil {
		fmt.Printf("Error happened while reading: %v\n", intentionalErr)
	}

	readArticleResponse, err := client.ReadArticle(context.Background(), &blogpb.ReadArticleRequest{ArticleId: articleId})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	fmt.Printf("Article found: %v\n", readArticleResponse)

	// update article
	newArticle := &blogpb.Article{
		Id:       articleId,
		AuthorId: "Another Author",
		Title:    "First article edited",
		Content:  "Content with additions",
	}

	updateArticleResponse, err := client.UpdateArticle(context.Background(), &blogpb.UpdateArticleRequest{Article: newArticle})
	if err != nil {
		fmt.Printf("Error happened while updating: %v \n", err)
	}
	fmt.Printf("Article was updated: %v\n", updateArticleResponse)

	// Delete article

	deleteResponse, deleteErr := client.DeleteArticle(context.Background(), &blogpb.DeleteArticleRequest{
		ArticleId: articleId,
	})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", err)
	}

	fmt.Printf("Article was deleted. %v\n", deleteResponse)

	// List all articles

	stream, err := client.ListArticles(context.Background(), &blogpb.ListArticlesRequest{})
	if err != nil {
		log.Fatalf("Error while calling ListArticles RPC: %v", err)
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
		fmt.Println(response.GetArticle())
	}
}
