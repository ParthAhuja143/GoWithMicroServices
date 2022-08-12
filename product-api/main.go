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
	HTTPmiddleware "github.com/ParthAhuja143/GoWithMicroServices/middlewares"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()
	logger := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handlers
	productsHandler := handlers.NewProducts(logger)

	// create a new serve mux and register the handlers
	//serveMux := http.NewServeMux()
	serveMux := mux.NewRouter()
	
	
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandler.GetProducts)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productsHandler.AddProduct)
	postRouter.Use(HTTPmiddleware.MiddlewareValidateProduct)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(HTTPmiddleware.MiddlewareValidateProduct)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.DeleteProduct)
	
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	
	//serveMux.("/products/", productsHandler)

	server := http.Server{
		Addr: *bindAddress,
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