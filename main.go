package main

import (
	"fmt"
	"interview-accountapi-demo/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()     //create a new router instance
	client := &http.Client{} //create a new http client

	//register endpoints with their corresponding handlers
	r.HandleFunc("/organisation/accounts", handlers.CreateHandler(client)).Methods(http.MethodPost)
	r.HandleFunc("/organisation/accounts", handlers.GetHandler(client)).Methods(http.MethodGet)
	r.HandleFunc("/organisation/accounts", handlers.DeleteHandler(client)).Methods(http.MethodDelete)

	fmt.Println("Starting server at port 8000...")
	http.ListenAndServe(":8000", r) //start listening on the port

}
