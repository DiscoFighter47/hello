package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/gojektech/heimdall/hystrix"
)

const url = "http://10.83.4.216/hello/"

var tPath = "./views/hello.html"

var t *template.Template
var client *hystrix.Client

func init() {
	t, _ = template.ParseFiles(tPath)
	client = hystrix.NewClient(
		hystrix.WithHTTPTimeout(10*time.Second),
		hystrix.WithCommandName("hello server request"),
		hystrix.WithHystrixTimeout(10*time.Second),
		hystrix.WithMaxConcurrentRequests(1000),
		hystrix.WithErrorPercentThreshold(20),
	)
}

func call(name string) string {
	resp, err := client.Get(url+name, nil)
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.EscapedPath())
	if name == "hello" {
		name = "world"
	}
	log.Println("request from ", name)
	t.Execute(w, call(name))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	log.Println("starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("shutting down server...")
		log.Fatal(err)
	}
}
