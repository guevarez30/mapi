package main

import (
	"fmt"
	"guevarez30/mapi/pkg/database"
	"guevarez30/mapi/pkg/order"
	"guevarez30/mapi/pkg/queue"
	"guevarez30/mapi/pkg/workgroup"
)

func main() {
	MAX_WORKERS := 3

	jobs := make(chan order.Order, 3)
	quitConsume := make(chan bool, 1)
	results := make(chan workgroup.Result)

	db := database.InitDb()

	wg := workgroup.WorkGroup{
		MaxWorkers: MAX_WORKERS,
		Jobs:       jobs,
		Results:    results,
		DB:         db,
	}

	wg.Run()
	queue.Subscribe(jobs, quitConsume)

	for {
		select {
		case res := <-results:
			fmt.Println(res)
		}
	}
}
