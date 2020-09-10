package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var iterations int = 100

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Average Time To Respond
func responseTimes(url string) []time.Duration {
	timePerRequest := make(chan time.Duration, iterations)
	resTimes := make([]time.Duration, iterations)
	for i := 0; i < iterations; i++ {
		go timeToRespond(url, &timePerRequest)
		time.Sleep(250 * time.Millisecond)
	}
	for i := 0; i < iterations; i++ {
		select {
		case tempTIme := <-timePerRequest:
			resTimes = append(resTimes, tempTIme)
		}
	}
	return resTimes
}

func getCommandLineInput() string {
	url := flag.String("url", "https://docs.google.com/", "Url to perform health check")
	flag.Parse()
	return *url
}

func main() {
	// url := getCommandLineInput()
	// averageTime := float64(responseTimes(url)) / (float64(iterations) * 1e9)
	// fmt.Println(averageTime)
}
