package main

import (
	"GranicnaPolicija/db"
	"GranicnaPolicija/handlers"
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
		port = "8083"
	}

	// NoSQL: Initialize Product Repository store
	repo, err := db.NewGPRepoDB(timeoutContext, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer repo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	repo.Ping()

	//Initialize the handler and inject said logger
	policajacHandler := handlers.NewPolicajacHandler(logger, repo)
	prelazakHandler := handlers.NewPrelazakHandler(logger, repo)
	proveraHandler := handlers.NewProveraHandler(logger, repo)

	//Initialize the router and add a middleware for all the requests
	routerUser := mux.NewRouter()
	routerUser.Use(policajacHandler.MiddlewareContentTypeSet)

	//Granicni policajac routers
	postPolicajacRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postPolicajacRouter.HandleFunc("/gpolicajac", policajacHandler.Register)
	postPolicajacRouter.Use(policajacHandler.MiddlewareGPValidation)

	loginPolicajacRouter := routerUser.Methods(http.MethodPost).Subrouter()
	loginPolicajacRouter.HandleFunc("/", policajacHandler.LoginPolicajac)

	//1Y  Provera lične karte i pasoša sa MUP-om (Mup nije uradio pasose tkd proveravam samo licnu kartu)
	//2Y  Provera postojanja poternica lica sa MUP-om
	//3Y  Zabrana prelaska licu u slucaju utvrdjenog pokusaja krijumcarenja i pisanje krivicne prijave
	//4Y  Zabrana prelaska licu sa poternicom i pisanje krivicne prijave
	//4Y  Pustanje lica preko granice tj kreiranje prelaska
	//5Y  Lista svih prelaza granice do sada
	//6Y  Prosledjivanje krivicnih prijava tužilaštvu (lista krivicnih prijava gde se prosledjuje preko fronte i menja se status na prosledjeno)

	//TODO : na frontu staviti dropdown menu za tuzilastva radi cuvanja tj slanja krivicne prijave

	//TODO : svaka provera na cekanju ima accept i refuse dugme... accept radi se stavka 1A, refuse onda stavka 2B
	//TODO : na frontu ofarbati u crvenu provere sa poternicom i blokirati dugme accept
	//TODO : 1A na frontu kada se prihvata prelazak mora da se kreira prelazak i ako vrati ok salje se accept provera
	//TODO : 2B na frontu kada se odbije prelazak prebacuje se na pisanje krivicne prijave i samim tim se salje tuzilastvu

	//Provera Gradjana routers
	postProveraRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postProveraRouter.HandleFunc("/provera", proveraHandler.CreateProveraHandler)
	//postProveraRouter.Use(proveraHandler.MiddlewareProveraValidation)

	getProvereRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getProvereRouter.HandleFunc("/provere", proveraHandler.GetProvere)

	getProvereNaCekanjuRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getProvereNaCekanjuRouter.HandleFunc("/provere/cekaju", proveraHandler.GetProvereNaCekanju)

	postAcceptProvera := routerUser.Methods(http.MethodPost).Subrouter()
	postAcceptProvera.HandleFunc("/provera/accept/{proveraId}", proveraHandler.AcceptPrelazak)

	postRefuseProvera := routerUser.Methods(http.MethodPost).Subrouter()
	postRefuseProvera.HandleFunc("/provera/refuse/{proveraId}", proveraHandler.RefusePrelazak)

	//Prelazak granice routers
	postPrelazakRouter := routerUser.Methods(http.MethodPost).Subrouter()
	postPrelazakRouter.HandleFunc("/prelazak", prelazakHandler.CreatePrelazakHandler)
	//postPrelazakRouter.Use(prelazakHandler.MiddlewarePrelazakValidation)

	getPrelazakRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPrelazakRouter.HandleFunc("/prelasci", prelazakHandler.GetPrelasci)

	getPrelazakByGranicaRouter := routerUser.Methods(http.MethodGet).Subrouter()
	getPrelazakByGranicaRouter.HandleFunc("/prelasci/{g_prelaz}", prelazakHandler.GetPrelasci)

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
