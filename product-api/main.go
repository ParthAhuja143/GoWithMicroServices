package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ParthAhuja143/GoWithMicroServices/handlers"
)

func main() {

	logger := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handlers
	productsHandler := handlers.NewProducts(logger)

	// create a new serve mux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/products/", productsHandler)

	server := http.Server{
		Addr: ":9090",
		Handler: serveMux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	// It's in a go routine so that it doesn't stop execution of code below it
	go func(){
		logger.Println("Server running on port 9090")
		err := server.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	// We make a channel to communicate and we only listen for Interrupt or Kill
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	// This execution is blocked until we get a message(which is Kill or Interrupt) and then graceful shutdown begins
	signal := <-signalChannel
	logger.Println("Received terminate, graceful shutdown", signal)

	serverTimeoutContext, cancelContextFunc := context.WithTimeout(context.Background(), 30*time.Second)
	
	defer cancelContextFunc()

	server.Shutdown(serverTimeoutContext)
}