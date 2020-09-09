package main

import (
	"fmt"
	"net/http"
	"time"
)

// Time To Respond
func timeToRespond(url string, t *chan time.Duration) {
	startTime := time.Now()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(res.StatusCode)
		panic(err)
	}
	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
	}
	defer res.Body.Close()
	diff := time.Now().Sub(startTime)
	*t <- diff
}

// Average Time To Respond
func averageTimeToRespond(url string, iterations int) time.Duration {
	timePerRequest := make(chan time.Duration, iterations)
	initialTime := time.Duration(0)
	for i := 0; i < iterations; i++ {
		go timeToRespond(url, &timePerRequest)
	}
	for i := 0; i < iterations; i++ {
		select {
		case tempTIme := <-timePerRequest:
			initialTime += tempTIme
		}
	}
	return initialTime
}

func main() {
	averageTime := float64(averageTimeToRespond("https://docs.google.com/", 5)) / (5 * 1e9)
	fmt.Println(averageTime)
}
