package main

import (
	"log"
	"net/http"
	"path"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.EscapedPath())
	if name == "hello" {
		name = "world"
	}
	log.Println("request from ", name)
	time.Sleep(5 * time.Second)
	w.Write([]byte("hello " + name + "!"))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	log.Println("starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("shutting down server...")
		log.Fatal(err)
	}
}
