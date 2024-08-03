package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello in golang")
	r:= mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("get")
	log.Fatal(http.ListenAndServe(":3000",r))
}
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}
