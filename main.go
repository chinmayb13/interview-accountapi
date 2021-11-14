package main

import (
	"fmt"
	"interview-accountapi-demo/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	client := &http.Client{}
	//sr := r.PathPrefix("https://api.staging-form3.tech/v1").Subrouter()
	r.HandleFunc("/organisation/accounts", handlers.CreateHandler(client)).Methods(http.MethodPost)
	r.HandleFunc("/organisation/accounts", handlers.GetHandler(client)).Methods(http.MethodGet)
	r.HandleFunc("/organisation/accounts", handlers.DeleteHandler(client)).Methods(http.MethodDelete)
	fmt.Println("Starting server at port 8000...")
	http.ListenAndServe(":8000", r)

}
