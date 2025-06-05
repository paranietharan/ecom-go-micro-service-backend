package main

import (
	"ecom-go-micro-service-backend/ecom-api/handler"
	"ecom-go-micro-service-backend/ecom-grpc/pb"
	"ecom-go-micro-service-backend/env"
	"log"

	"github.com/ianschenck/envflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const minSecretKeySize = 32

func main() {
	err := env.LoadEnv()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	var (
		secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")
		svcAddr   = envflag.String("GRPC_SVC_ADDR", "0.0.0.0:9091", "address where the ecomm-grpc service is listening on")
	)

	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("SECRET_KEY must be at least %d characters", minSecretKeySize)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(*svcAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewEcommClient(conn)
	hdl := handler.NewHandler(client, *secretKey)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
