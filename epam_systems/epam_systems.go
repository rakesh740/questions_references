package main

import (
	"fmt"
	"sync"
)


// create a new web server - and rate limiter
// create a dynamic worker pool. where number of workers can be increased by load amount

// func main() {
// 	fmt.Println("Hello World")

// 	r := gin.Default()

// 	r.GET("/ping", Index)

// 	r.Run(":8080")
// }


// func Index(c *gin.Context) {
// 	//
// 	c.JSON(
// 		200,
// 		gin.H{"message": "pong"},
// 	)
// }

// basic worker - pool - 3 worker -

func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	defer close(out)
	for v := range in {
		fmt.Printf("workerID: %v, incoming %v\n", workerID, v)
		out <- v * v
	}
}



type Pool struct {
	workerNum int
	data      chan<- int // use buffered channel - check if sending is blocked
	//
	// time
}

func main() {

	var numWorker int = 3

	data := make(chan int)
	response := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(data chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(data)
		for i := 0; i < 10; i++ {
			data <- i
		}
	}(data, &wg)

	for i := 0; i < numWorker; i++ {
		go worker(data, response, &wg, i)
	}

	go func(response <-chan int, wg *sync.WaitGroup) {
		for v := range response {
			fmt.Printf("v: %v\n", v)
		}
	}(response, &wg)

	for {
		select {
		case <-data:

		}
	}

	wg.Wait()

	p := paypal{}

	p.getPayeeName()

}

type IPayment interface {
	pay()
}

type bitcoin struct {
	baseUser
}

func (b *bitcoin) pay() {

}

type paypal struct {
	baseUser
}

func (p paypal) pay() {

}

type baseUser struct {
	name string
}

func (u baseUser) getPayeeName() string, error {
	return u.name
}

// ID , log_time, log_type, log_message - sql query for top 5 repeated log message 
select count(ID) , log_time, log_type, log_message from log_table group by log_message order by count(ID) desc limit 5;

// get second highest salary for employee from employee table