package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	listenAddr := "8080"

	// create a logger, router and server
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	router := newRouter()
	middlewareList := middlewares{
		logging(logger),
		tracing(nextRequestIdCallback),
	}
	server := newServer(
		listenAddr,
		middlewareList.apply(router),
		logger,
	)

	// run our server
	if err := server.run(); err != nil {
		log.Fatal(err)
	}
}

func nextRequestIdCallback() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
