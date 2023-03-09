package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("Fatal error occurred in server setup: ", err)
	}
}

func run() error {
	addr := "0.0.0.0:9090"
	srv := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(reqHandler),
	}

	log.Println("Listening on: ", addr)
	return srv.ListenAndServe()
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf(`Welcome to COLORMATIC`)
	_, err := w.Write([]byte(body))
	if err != nil {
		log.Println("Failed to write response: ", err)
	}
}
