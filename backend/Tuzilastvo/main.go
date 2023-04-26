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
	usersHandler := handlers.NewHandler(logger, repo)

	//Initialize the router and add a middleware for all the requests
	routerUser := mux.NewRouter()
	routerUser.Use(usersHandler.MiddlewareContentTypeSet)

	postTuzilastvoRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postTuzilastvoRouter.HandleFunc("/tuzilastva", usersHandler.CreateTuzilastvo)
	postTuzilastvoRouter.Use(usersHandler.MiddlewareTuzilastvoValidation)

	getTuzilastvaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getTuzilastvaRouter.HandleFunc("/tuzilastva", usersHandler.GetTuzilastva)

	getTuzilastvoRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getTuzilastvoRouter.HandleFunc("/tuzilastva/{id}", usersHandler.GetTuzilastvo)

	getPrijaveRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPrijaveRouter.HandleFunc("/prijave", usersHandler.GetPrijave)

	getPrijavaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPrijavaRouter.HandleFunc("/prijave/{id}", usersHandler.GetPrijava)

	postPrijavaRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postPrijavaRouter.HandleFunc("/prijave", usersHandler.CreatePrijava)
	postPrijavaRouter.Use(usersHandler.MiddlewarePrijavaValidation)

	confirmPrijavaRouter := routerUser.Methods(http.MethodPut).Subrouter()
	confirmPrijavaRouter.HandleFunc("/prijave/confirm/{id}", usersHandler.ConfirmPrijava)
	confirmPrijavaRouter.Use(usersHandler.MiddlewarePrijavaValidation)

	declinePrijavaRouter := routerUser.Methods(http.MethodPut).Subrouter()
	declinePrijavaRouter.HandleFunc("/prijave/decline/{id}", usersHandler.DeclinePrijava)
	declinePrijavaRouter.Use(usersHandler.MiddlewarePrijavaValidation)

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
