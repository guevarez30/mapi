package workers

import (
	"fmt"
	Order "guevarez30/mapi/pkg/order"
	"time"
)

type Result struct {
	Status string
}

type WorkerGroup struct {
	MaxWorkers int
	jobs       <-chan Order.Order
	results    chan<- Result
}

func New(maxWorkers int, jobs <-chan Order.Order, results chan<- Result) WorkerGroup {
	return WorkerGroup{
		MaxWorkers: maxWorkers,
		jobs:       jobs,
		results:    results,
	}
}

func sleep(wId int) {
	sleepy := wId * wId
	time.Sleep(time.Duration(sleepy) * time.Second)
}

func handler(wId int, jobs <-chan Order.Order, results chan<- Result) {
	for {
		select {
		default:

			fmt.Printf("Worker: %d Doing Order\n", wId)
			<-jobs
			sleep(wId)
			fmt.Printf("Worker: %d Completed\n", wId)
		}
	}

}

func (wg WorkerGroup) Run() {
	for w := 1; w <= wg.MaxWorkers; w++ {
		fmt.Println("Creating workers")
		go handler(w, wg.jobs, wg.results)
	}
}
