package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sharansharma94/nick/product-api/handlers"
)

func main() {
	fmt.Println("Hello")

	ll := log.New(os.Stdout, "product", log.LstdFlags)
	hh := handlers.NewHello(ll)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", hh)

	server := http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		ErrorLog:     ll,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		ll.Println("server started on port 8080")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c

	log.Println("got signal : ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}
