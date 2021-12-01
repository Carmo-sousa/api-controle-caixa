package main

import (
	"fmt"
	"time"
)

func main() {
	requesId := make(chan int)
	concurrency := 3

	for i := 1; i <= concurrency; i++ {
		go work(requesId, i)
	}

	for i := 1; i <= 100; i++ {
		requesId <- i
	}
}

func work(requestId chan int, worker int) {
	for req := range requestId {
		// res, err := http.Get("http://localhost:5000/api/v1/users/")

		// if err != nil {
		// 	log.Fatal(err.Error())
		// }

		// defer res.Body.Close()

		// body, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	log.Fatal(err.Error())
		// }

		fmt.Printf("Worker: %v RequestId: %v\n", worker, req)
		time.Sleep(time.Second)
	}
}
