package main

import (
	"fmt"
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

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", HelloGet).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadGet).Methods(http.MethodGet)

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
