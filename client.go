package main

import (
	"context"
	pb "golang-grpc-bcrypt-service/golanggrpcbcryptservice"
	"log"

	"google.golang.org/grpc"
)

func generateBCrypt(client pb.BCryptClient, password *pb.PasswordRequest) {
	feature, err := client.PostBCrypt(context.Background(), password)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:6565", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewBCryptClient(conn)

	generateBCrypt(client, &pb.PasswordRequest{Password: "password"})
}
