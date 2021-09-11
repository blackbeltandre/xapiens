package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"xapiens/app"
	"xapiens/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/admin/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/v1/admin/login", controllers.Authenticate).Methods("POST")
	
	router.HandleFunc("/api/v1/user/new", controllers.CreateUsers).Methods("POST")
	router.HandleFunc("/api/v1/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/detail", controllers.GetUsersDetail).Methods("GET")
	router.HandleFunc("/api/v1/user/delete", controllers.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/api/v1/user/update", controllers.UpdateUsers).Methods("PATCH")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("port")
	if port == "" {
		port = "8002" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8002/api
	if err != nil {
		fmt.Print(err)
	}
}
