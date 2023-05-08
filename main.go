package main

import (
	"fmt"
	Order "guevarez30/mapi/pkg/order"
	"guevarez30/mapi/pkg/queue"
	"guevarez30/mapi/pkg/workers"
	"time"
)

func queueJob(jobs chan<- Order.Order) {
	url := "https://google.com"
	jobs <- Order.Order{URL: url}
	time.Sleep(3)
	queueJob(jobs)
}

func main() {
	fmt.Println("Starting consuming")
	jobs := make(chan Order.Order, 3)
	results := make(chan workers.Result)

	// Mock to produce message on job queue
	go queueJob(jobs)

	wg := workers.New(3, jobs, results)
	wg.Run()
	go queue.Subscribe(jobs)

	for {
		select {
		case res := <-results:
			fmt.Println(res)
		}
	}

}
