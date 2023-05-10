package queue

import (
	"fmt"
	"guevarez30/mapi/pkg/order"
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

// Keep Calling Subscribe to read next message
func Subscribe(jobs chan<- order.Order, quit <-chan bool) {
	go func() {
		for {
			select {
			case <-quit:
				return
			default:

				o := order.Order{
					Task:     "photogrammatry",
					Details:  "banana",
					UserUUID: "b5c6379a-ebf9-4845-841b-e187ece03d4d",
				}

				jobs <- o
				// Dont over board the queue
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
