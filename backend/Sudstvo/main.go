package main

import (
	"Sudstvo/db"
	"Sudstvo/handlers"
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

	logger := log.New(os.Stdout, "[sudstvo-api] ", log.LstdFlags)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	port := os.Getenv("app_port")
	if len(port) == 0 {
		port = "8082"
	}

	// NoSQL: Initialize Product Repository store
	repo, err := db.SudstvoRepoDB(timeoutContext, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer repo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	repo.Ping()

	//Initialize the handler and inject said logger
	poternicaHandler := handlers.NewPoternicaHandler(logger, repo)
	sudHandler := handlers.NewSudHandler(logger, repo)
	rocisteHandler := handlers.NewRocisteHandler(logger, repo)
	sudijaHandler := handlers.NewSudijaHandler(logger, repo)

	//Initialize the router and add a middleware for all the requests
	routerUser := mux.NewRouter()

	postSudRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postSudRouter.HandleFunc("/sudovi", sudHandler.CreateSud)
	postSudRouter.Use(sudHandler.MiddlewareSudValidation)

	getSudoviRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getSudoviRouter.HandleFunc("/sudovi", sudHandler.GetSudovi)

	getSudRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getSudRouter.HandleFunc("/sudovi/{id}", sudHandler.GetSud)

	getPoterniceRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPoterniceRouter.HandleFunc("/poternice", poternicaHandler.GetPoternice)

	getPoternicaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPoternicaRouter.HandleFunc("/poternice/{id}", poternicaHandler.GetPoternica)

	postPoternicaRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postPoternicaRouter.HandleFunc("/poternice", poternicaHandler.CreatePoternica)
	postPoternicaRouter.Use(poternicaHandler.MiddlewarePoternicaValidation)

	postRocisteRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postRocisteRouter.HandleFunc("/rocista", rocisteHandler.CreateRociste)
	postRocisteRouter.Use(rocisteHandler.MiddlewareRocisteValidation)

	getRocistaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getRocistaRouter.HandleFunc("/rocista", rocisteHandler.GetRocista)

	getRocisteRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getRocisteRouter.HandleFunc("/rocista/{id}", rocisteHandler.GetRocista)

	loginRouter := routerUser.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", sudijaHandler.LoginUser)

	registerRouter := routerUser.Methods(http.MethodPost).Subrouter()
	registerRouter.HandleFunc("/register", sudijaHandler.Register)
	registerRouter.Use(sudijaHandler.MiddlewareSudijaValidation)

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
