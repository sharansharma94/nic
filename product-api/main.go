package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello handler")

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error in reading input ", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "hello %s", data)
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
