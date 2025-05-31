package main

import (
	"ecom-go-micro-service-backend/db"
	"ecom-go-micro-service-backend/ecom-api/handler"
	"ecom-go-micro-service-backend/ecom-api/server"
	"ecom-go-micro-service-backend/ecom-api/storer"
	"fmt"
	"log"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("database connection failed.......")
	}
	defer db.Close()

	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRoutes(hdl)

	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Sucessfully running in 8080 ...........")
	}
}
