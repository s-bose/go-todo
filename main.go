package main

import (
	"go-todo/app/controllers"
	"go-todo/db"
	"log"
	"net/http"
)

func main() {

	database, err := db.InitDatabase()
	if err != nil {
		log.Fatal("something went wrong while setting up database", err)
	}

	routerV1 := http.NewServeMux()

	routerV1.Handle("/user", controllers.CreateUserRouter(database))

	mux := http.NewServeMux()
	mux.Handle("/api/v1", routerV1)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening on port :8080")
	server.ListenAndServe()
}
