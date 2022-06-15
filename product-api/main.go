package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "GO Microservice", log.LstdFlags)
	helloHandler := handlers.(logger)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":9000", serveMux)
}