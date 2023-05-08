package workers

import (
	"fmt"
	Order "guevarez30/mapi/pkg/order"
	"log"
	"math/rand"
	"net/http"
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
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5) // n will be between 0 and 10
	n = n + 5
	fmt.Printf("Worker: %d Sleeping %d seconds...\n", wId, n)
	time.Sleep(time.Duration(n) * time.Second)
}

func handler(wId int, jobs <-chan Order.Order, results chan<- Result) {
	for order := range jobs {
		sleep(wId)
		_, err := http.Get(order.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: fmt.Sprintf("Worker: %d finished", wId)}
	}
}

func (wg WorkerGroup) Run() {
	for w := 1; w <= wg.MaxWorkers; w++ {
		go handler(w, wg.jobs, wg.results)
	}
}
