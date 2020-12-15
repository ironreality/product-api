package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	addRouter := sm.Methods(http.MethodPost).Subrouter()
	addRouter.HandleFunc("/", ph.AddProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	s := &http.Server{
		Addr:         ":9000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Fatal(err)
		}
	}()
	log.Print("Server started")

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := s.Shutdown(tc)
	if err != nil {
		log.Fatalf("Server shutdown Failed: %v", err)
	}
	log.Print("Server exited properly")
}
