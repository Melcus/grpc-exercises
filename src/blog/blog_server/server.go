package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"synapse/src/blog/blogpb"
	"time"
)

var articlesCollection *mongo.Collection

type Article struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type server struct{}

func (*server) CreateArticle(ctx context.Context, request *blogpb.CreateArticleRequest) (*blogpb.CreateArticleResponse, error) {
	fmt.Println("Create article request")
	article := request.GetArticle()

	data := Article{
		AuthorId: article.GetAuthorId(),
		Content:  article.GetContent(),
		Title:    article.GetTitle(),
	}

	result, err := articlesCollection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to ObjectId"))
	}

	return &blogpb.CreateArticleResponse{
		Article: &blogpb.Article{
			Id:       objectId.Hex(),
			AuthorId: article.GetAuthorId(),
			Title:    article.GetTitle(),
			Content:  article.GetContent(),
		},
	}, nil
}

func (*server) ReadArticle(ctx context.Context, request *blogpb.ReadArticleRequest) (*blogpb.ReadArticleResponse, error) {
	fmt.Println("Read article request")
	objectId, err := primitive.ObjectIDFromHex(request.GetArticleId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	data := &Article{}

	result := articlesCollection.FindOne(context.Background(), bson.M{"_id": objectId})
	if err := result.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find Article with specified ID: %v", err))
	}
	return &blogpb.ReadArticleResponse{
		Article: dataToArticlePB(data),
	}, nil
}

func (*server) UpdateArticle(ctx context.Context, request *blogpb.UpdateArticleRequest) (*blogpb.UpdateArticleResponse, error) {
	fmt.Println("Update Article Request")
	article := request.GetArticle()

	objectId, err := primitive.ObjectIDFromHex(article.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	data := &Article{}

	result := articlesCollection.FindOne(context.Background(), bson.M{"_id": objectId})
	if err := result.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find Article with specified ID: %v", err))
	}

	data.AuthorId = article.GetAuthorId()
	data.Content = article.GetContent()
	data.Title = article.GetTitle()

	_, updateErr := articlesCollection.ReplaceOne(context.Background(), bson.M{"_id": objectId}, data)
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot update object in MongoDB: %v\n", updateErr))
	}

	return &blogpb.UpdateArticleResponse{
		Article: dataToArticlePB(data),
	}, nil
}

func (*server) DeleteArticle(ctx context.Context, request *blogpb.DeleteArticleRequest) (*blogpb.DeleteArticleResponse, error) {
	fmt.Println("Delete Article Request")

	objectId, err := primitive.ObjectIDFromHex(request.GetArticleId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	deleteResponse, deleteErr := articlesCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot delete object from MongoDB: %v\n", deleteErr))
	}
	if deleteResponse.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find object in MongoDB"))
	}

	return &blogpb.DeleteArticleResponse{
		Success: true,
	}, nil
}

func (*server) ListArticles(request *blogpb.ListArticlesRequest, stream blogpb.ArticleService_ListArticlesServer) error {
	fmt.Println("List Articles Request")

	cursor, err := articlesCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error %v \n", err))
	}
	defer func() {
		cursorCloseErr := cursor.Close(context.Background())
		if cursorCloseErr != nil {
			log.Fatalf("Failed to close MongoDB cursor")
		}
	}()

	for cursor.Next(context.Background()) {
		data := &Article{}
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB %v \n", err))
		}

		_ = stream.Send(&blogpb.ListArticlesResponse{
			Article: dataToArticlePB(data),
		})
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB %v \n", err))
	}

	return nil
}

func dataToArticlePB(data *Article) *blogpb.Article {
	return &blogpb.Article{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
	}
	fmt.Println("MongoDB ping OK ! - Blog Service started")

	articlesCollection = client.Database("blog_db").Collection("articles")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tls := false
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
	blogpb.RegisterArticleServiceServer(s, &server{})
	// Register reflection service on gRPC server
	reflection.Register(s)

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
	fmt.Println("Closing MongoDB Connection")
	_ = client.Disconnect(context.TODO())
	fmt.Println("End of program")
}
