package main

import (
	"ecom-go-micro-service-backend/db"
	"ecom-go-micro-service-backend/ecom-api/handler"
	"ecom-go-micro-service-backend/ecom-api/server"
	"ecom-go-micro-service-backend/ecom-api/storer"
	"fmt"
	"log"

	"github.com/ianschenck/envflag"
)

const minSecretKeySize = 32

func main() {
	var secretKey = envflag.String("SECRET_KEY", "01234567890123456789012345678901", "secret key for JWT signing")
	if len(*secretKey) < minSecretKeySize {
		log.Fatalf("SECRET_KEY must be at least %d characters", minSecretKeySize)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")

	// do something with the database
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv, *secretKey)
	handler.RegisterRoutes(hdl)
	fmt.Println("Starting ECOM API server on port 8080...")
	handler.Start(":8080")
}
