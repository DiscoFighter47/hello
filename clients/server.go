package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gojektech/heimdall/hystrix"
)

const path = "/hello/zahid"
const maxReq = 10000

var counter = 0
var lock sync.Mutex
var client *hystrix.Client

func init() {
	client = hystrix.NewClient(
		hystrix.WithHTTPTimeout(10*time.Second),
		hystrix.WithCommandName("hello server request"),
		hystrix.WithHystrixTimeout(10*time.Second),
		hystrix.WithMaxConcurrentRequests(1000),
		hystrix.WithErrorPercentThreshold(20),
	)
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
