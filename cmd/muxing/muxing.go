package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// HelloGet to handle Get with param
func HelloGet(w http.ResponseWriter, r *http.Request) {

	// read param from request
	param, _ := mux.Vars(r)["PARAM"]

	_, err := fmt.Fprintf(w, "Hello, %s!", param)
	if err != nil {
		fmt.Println(err)
		panic(err) // panic just for fun!)
	}
	w.WriteHeader(http.StatusOK)

}

// BadGet to handle Get with status 500
func BadGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}

// BodyPost to handle Post with data and param
func BodyPost(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading data")
	}
	data := []byte(fmt.Sprintf("I got message:\n%s", d))
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// HeadersPost to handle Post with params in header
func HeadersPost(w http.ResponseWriter, r *http.Request) {
	// read "a" and "b" from header
	a, b := r.Header.Get("a"), r.Header.Get("b")

	// convert them to string
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)

	// write sum to header
	w.Header().Set("a+b", strconv.Itoa(aInt+bInt))
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	// create handlers
	router.HandleFunc("/name/{PARAM}", HelloGet).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadGet).Methods(http.MethodGet)
	router.HandleFunc("/data", BodyPost).Methods(http.MethodPost)
	router.HandleFunc("/headers", HeadersPost).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
