package queue

import (
	"fmt"
	Order "guevarez30/mapi/pkg/order"
	"math/rand"
	"time"
)

func sleep(wId int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5) // n will be between 0 and 10
	n = n + 5
	fmt.Printf("Worker: %d Sleeping %d seconds...\n", wId, n)
	time.Sleep(time.Duration(n) * time.Second)
}

// Keep Calling Subscribe to artifically put orders on queue
func Subscribe(jobs chan<- Order.Order) {
	url := "https://google.com"
	jobs <- Order.Order{URL: url}
	time.Sleep(3)
	Subscribe(jobs)
}
