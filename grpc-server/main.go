package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "grpc-server/advice"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedAdviceServer
}

type SlipDTO struct {
	ID     int    `json:"id"`
	Advice string `json:"advice"`
}

type AdviceDTO struct {
	Slip SlipDTO `json:"slip"`
}

func (s *server) GetAdvice(ctx context.Context, in *pb.Empty) (*pb.AdviceMessage, error) {
	advice := ""

	res, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		advice = "I'm sorry, I cannot provide you with advice at the moment."
		return &pb.AdviceMessage{Message: advice}, nil
	}

	defer res.Body.Close()

	var data AdviceDTO

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		advice = "I'm sorry, I cannot provide you with advice at the moment."
	} else {
		advice = data.Slip.Advice
	}

	return &pb.AdviceMessage{Message: advice}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdviceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
