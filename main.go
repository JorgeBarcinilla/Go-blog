package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-blog/controllers"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Server started on: http://localhost:8000")

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/show", controllers.Show).Methods("GET")
	router.HandleFunc("/new", controllers.New).Methods("GET")
	router.HandleFunc("/edit", controllers.Edit).Methods("GET")
	router.HandleFunc("/insert", controllers.Insert).Methods("POST")
	router.HandleFunc("/update", controllers.Update).Methods("POST")
	router.HandleFunc("/delete", controllers.Delete).Methods("GET")
	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
