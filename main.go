package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	//jwt "github.com/dgrijalva/jwt-go"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi!")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Incremented")
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	fmt.Println("Starting server")

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", echoString)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
