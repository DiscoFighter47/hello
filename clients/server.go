package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

const path = "/hello/zahid"
const maxReq = 5000

var counter = 0
var lock sync.Mutex
var client *httpclient.Client

func init() {
	timeout := 15 * time.Second
	client = httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
}

func print(val int) {
	lock.Lock()
	counter += val
	log.Println("request count: ", counter)
	lock.Unlock()
}

func call(wg *sync.WaitGroup, url string) {
	print(1)
	_, err := client.Get(url, nil)
	if err != nil {
		log.Println(err)
	} else {
		print(-1)
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < maxReq; i++ {
		wg.Add(1)
		go call(&wg, "http://"+os.Args[1]+path)
	}
	wg.Wait()
	log.Println("request served: ", maxReq-counter)
}
