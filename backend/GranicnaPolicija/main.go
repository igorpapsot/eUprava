package main

import (
	"context"
	"log"
	"os"
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
}
