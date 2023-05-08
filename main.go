package main

import (
	"fmt"
	Order "guevarez30/mapi/pkg/order"
	"guevarez30/mapi/pkg/queue"
	"guevarez30/mapi/pkg/workers"
)

func main() {
	fmt.Println("Starting consuming")
	jobs := make(chan Order.Order, 3)
	quitConsume := make(chan bool, 1)
	results := make(chan workers.Result)

	wg := workers.New(3, jobs, results)
	wg.Run()
	queue.Subscribe(jobs, quitConsume)

	for {
		select {
		case res := <-results:
			fmt.Println(res)
		}
	}

}
