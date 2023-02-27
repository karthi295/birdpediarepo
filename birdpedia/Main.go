package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthik/birdpedia/handler"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", handler.GetBirdHandler).Methods("GET")
	r.HandleFunc("/bird", handler.CreateBirdHandler).Methods("POST")
	return r
}

func main() {

	r := newRouter()
	//r.HandleFunc("/hello", handler).Methods("GET")
	//http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World!")
}
