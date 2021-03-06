package main

import (
	"fmt"
	"log"
	"net/http"
)

var numberOfRequests = 100 //total number of requests allowed
var timeLimit = 3600.0     //time limit in seconds

func main() {
	fmt.Println("Application has been started on : 8050")
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8050", rateLimit(nil, numberOfRequests, timeLimit)))
}

//Index : hello world example to make GET request to
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string("Hello World example for GET request."))
}
