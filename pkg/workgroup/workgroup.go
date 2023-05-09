package workgroup

import (
	"fmt"
	"guevarez30/mapi/pkg/order"
	"time"

	"gorm.io/gorm"
)

func sleep(wId int) {
	sleepy := wId * wId
	time.Sleep(time.Duration(sleepy) * time.Second)
}

type Result struct {
	Status string
}

type WorkGroup struct {
	MaxWorkers int
	Jobs       <-chan order.Order
	Results    chan<- Result
	DB         *gorm.DB
}

func (wg WorkGroup) handle(wId int) {
	for {
		select {
		default:
			fmt.Printf("Worker: %d Doing Order\n", wId)
			o := <-wg.Jobs
			order.Insert(wg.DB, &o)
			sleep(wId)
			res := Result{
				Status: fmt.Sprintf("Worker: %d Completed", wId),
			}
			wg.Results <- res
		}
	}
}

func (wg WorkGroup) Run() {
	for w := 1; w <= wg.MaxWorkers; w++ {
		fmt.Println("Creating workers")
		go wg.handle(w)
	}
}
