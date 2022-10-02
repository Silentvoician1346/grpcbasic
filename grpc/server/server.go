package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/basictest01/grpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

const mongodbURI = "mongodb://localhost:27017"

type server struct {
}

func (*server) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	blog := req.GetCountry()

	args := countryItem{
		code:      blog.GetCode(),
		name:      blog.GetName(),
		region:    blog.GetRegion(),
		subRegion: blog.GetSubRegion(),
	}

	res, err := collection.InsertOne(context.Background(), args)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("failed to insert one to database: %v", err),
		)
	}
	objectID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("failed to create objectID: %v", err),
		)
	}

	return &pb.CreateCountryResponse{
		Country: &pb.Country{
			Id:        objectID.Hex(),
			Code:      req.Country.GetCode(),
			Name:      req.Country.GetName(),
			Region:    req.Country.GetRegion(),
			SubRegion: req.Country.GetSubRegion(),
		},
	}, nil
}

type countryItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	code      string             `bson:"code"`
	name      string             `bson:"name"`
	region    string             `bson:"region"`
	subRegion string             `bson:"sub_region"`
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("starting basic service...")

	// connect to mongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb:", mongodbURI)

	collection = client.Database("basic_service").Collection("country")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	// generated_proto.RegisterBasicTestServer(s, server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server: %v", err)
		}
	}()

	// wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	fmt.Println("Basic Service server started")

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing MongoDB Connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of Program")
}
