package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "golang-grpc-bcrypt-service/golanggrpcbcryptservice"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

const (
	port = ":6565"
)

type golangBcryptServer struct{}

func (s *golangBcryptServer) PostBCrypt(ctx context.Context, password *pb.PasswordRequest) (*pb.PasswordHashedResponse, error) {
	hash := hashPassword(password.GetPassword())
	fmt.Printf("Hashed value: %v\n", string(hash))
	return &pb.PasswordHashedResponse{Hash: hash}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on port: %s", port)

	grpcServer := grpc.NewServer()
	pb.RegisterBCryptServer(grpcServer, &golangBcryptServer{})
	grpcServer.Serve(lis)
}

func hashPassword(password string) string {
	fmt.Printf("Value: %v\n", password)
	rounds, _ := strconv.Atoi(os.Getenv("ROUNDS"))
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), rounds)
	return string(hash)
}
