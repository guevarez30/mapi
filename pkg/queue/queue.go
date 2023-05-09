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
					Task:    "photogrammatry",
					Details: "banana",
					UserId:  "518031f7-bac1-43ba-b5fb-a6045b2e09de",
				}

				jobs <- o
				// Dont over board the queue
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
