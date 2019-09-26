package main

import (
	"fmt"
	"net/http"
	"os"

	"./controllers"

	"github.com/gorilla/mux"
)

func main() {
	// https://github.com/gorilla/mux
	// https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
	router := mux.NewRouter()

	router.HandleFunc("/api/getDetails/{table}/{id}", controllers.GetConnection).Methods("GET") //  user/2/contacts

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/
	if err != nil {
		fmt.Print(err)
	}
}
