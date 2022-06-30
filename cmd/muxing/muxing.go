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

func HelloGet(w http.ResponseWriter, r *http.Request) {

	param, _ := mux.Vars(r)["PARAM"]

	_, err := fmt.Fprintf(w, "Hello, %s", param)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func BadGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	//w.Header().Set("Content-Type", "application/json")
	//resp := make(map[string]string)
	//resp["message"] = "Some Error Occurred"
	//jsonResp, err := json.Marshal(resp)
	//if err != nil {
	//	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	//}
	//w.Write(jsonResp)
	return
}

func BodyPost(w http.ResponseWriter, r *http.Request) {
	//param, _ := mux.Vars(r)["PARAM"]
	d, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading data")
	}
	data := []byte(fmt.Sprintf("I got message:\n%s", d))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func HeadersPost(w http.ResponseWriter, r *http.Request) {
	a, b := r.Header.Get("a"), r.Header.Get("b")
	for k, v := range r.Header {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Print("aaa")
	fmt.Println(a)
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	w.Header().Set("a+b", strconv.Itoa(aInt+bInt))
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

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
