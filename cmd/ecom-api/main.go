package main

import (
	"ecom-go-micro-service-backend/db"
	"ecom-go-micro-service-backend/ecom-api/handler"
	"ecom-go-micro-service-backend/ecom-api/server"
	"ecom-go-micro-service-backend/ecom-api/storer"
	"ecom-go-micro-service-backend/env"
	"fmt"
	"log"
)

func main() {
	err := env.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
		log.Fatal("database connection failed.......")
	}
	defer db.Close()

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRoutes(hdl)
	fmt.Println("Starting ECOM API server on port 8080...")
	handler.Start(":8080")
}
