package main

import (
	"go-todo/app/controllers"
	"go-todo/app/db"
	"log"
	"net/http"
)

func main() {

	database, err := db.InitDatabase()
	if err != nil {
		log.Fatal("something went wrong while setting up database", err)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend service up and running"))
	})

	router.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", controllers.RegisterUserRoutes(database)))
	router.Handle("/api/v1/todos/", http.StripPrefix("/api/v1/todos", controllers.RegisterTodoRoutes(database)))

	http.ListenAndServe(":8000", router)

	defer database.CloseDatabase()
}
