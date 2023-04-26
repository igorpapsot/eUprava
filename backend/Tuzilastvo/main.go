package main

import (
	"Tuzilastvo/db"
	"Tuzilastvo/handlers"
	"context"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "[auth-api] ", log.LstdFlags)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	port := os.Getenv("app_port")
	if len(port) == 0 {
		port = "8080"
	}

	// NoSQL: Initialize Product Repository store
	repo, err := db.NewRepoDB(timeoutContext, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer repo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	repo.Ping()

	//Initialize the handler and inject said logger
	handler := handlers.NewHandler(logger, repo)

	//Initialize the router and add a middleware for all the requests
	routerUser := mux.NewRouter()
	routerUser.Use(handler.MiddlewareContentTypeSet)

	postTuzilastvoRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postTuzilastvoRouter.HandleFunc("/tuzilastva", handler.CreateTuzilastvo)
	postTuzilastvoRouter.Use(handler.MiddlewareTuzilastvoValidation)

	getTuzilastvaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getTuzilastvaRouter.HandleFunc("/tuzilastva", handler.GetTuzilastva)

	getTuzilastvoRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getTuzilastvoRouter.HandleFunc("/tuzilastva/{id}", handler.GetTuzilastvo)

	getPrijaveRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPrijaveRouter.HandleFunc("/prijave", handler.GetPrijave)

	getJavnePrijaveRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getJavnePrijaveRouter.HandleFunc("/prijave/public", handler.GetJavnePrijave)

	getPrijavaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPrijavaRouter.HandleFunc("/prijave/{id}", handler.GetPrijava)

	postPrijavaRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postPrijavaRouter.HandleFunc("/prijave", handler.CreatePrijava)
	postPrijavaRouter.Use(handler.MiddlewarePrijavaValidation)

	confirmPrijavaRouter := routerUser.Methods(http.MethodPut).Subrouter()
	confirmPrijavaRouter.HandleFunc("/prijave/confirm/{id}", handler.ConfirmPrijava)

	declinePrijavaRouter := routerUser.Methods(http.MethodPut).Subrouter()
	declinePrijavaRouter.HandleFunc("/prijave/decline/{id}", handler.DeclinePrijava)

	postOptuznicaRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postOptuznicaRouter.HandleFunc("/optuznice", handler.CreateOptuznica)
	postOptuznicaRouter.Use(handler.MiddlewareOptuznicaValidation)

	getOptuzniceRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getOptuzniceRouter.HandleFunc("/optuznice", handler.GetOptuznice)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"https://localhost:4200/"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,        // Addr optionally specifies the TCP address for the server to listen on, in the form "host:port". If empty, ":http" (port 80) is used.
		Handler:      cors(routerUser),  // handler to invoke, http.DefaultServeMux if nil
		IdleTimeout:  120 * time.Second, // IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
		ReadTimeout:  2 * time.Second,   // ReadTimeout is the maximum duration for reading the entire request, including the body. A zero or negative value means there will be no timeout.
		WriteTimeout: 4 * time.Second,   // WriteTimeout is the maximum duration before timing out writes of the response.
	}

	logger.Println("Server listening on port", port)

	//HTTPs:
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)
	signal.Notify(sigCh, syscall.SIGKILL)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)
	timeoutContext, _ = context.WithTimeout(context.Background(), 30*time.Second)

	//Try to shut down gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}

}
