package main

import (
	"fmt"
	"go-api-starter/routers"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	handler := routers.InitRouter()
	endpoint := fmt.Sprintf(":%s", port)

	server := &http.Server{
		Addr:    endpoint,
		Handler: handler,
	}

	log.Printf("Server Listening -  %s", endpoint)

	_ = server.ListenAndServe()
}
