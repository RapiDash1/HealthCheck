package check

import (
	"flag"
	"fmt"
	"healthCheck/pkg/bst"
	"net/http"
	"time"
)

var iterations int = 100
var sortedOrder []float64

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
func responseTimes(url string) {
	timePerRequest := make(chan time.Duration, iterations)
	for i := 0; i < iterations; i++ {
		go timeToRespond(url, &timePerRequest)
		time.Sleep(50 * time.Millisecond)
	}
	for i := 0; i < iterations; i++ {
		select {
		case tempTIme := <-timePerRequest:
			bst.AddNode(float64(tempTIme) / 1e9)
		}
	}
}

// Get Command Line Input
func getCommandLineInput() string {
	url := flag.String("url", "https://docs.google.com/", "Url to perform health check")
	flag.Parse()
	return *url
}

func Percentile99() {
	fmt.Printf("99th percentile: %v\n", sortedOrder[len(sortedOrder)-1])
}

func Percentile95() {
	fmt.Printf("95th percentile: %v\n", sortedOrder[len(sortedOrder)-6])
}

func Percentile90() {
	fmt.Printf("90th percentile: %v\n", sortedOrder[len(sortedOrder)-11])
}

func Percentile1() {
	fmt.Printf("1st percentile: %v\n", sortedOrder[0])
}

func PerformHealthCheck() {
	responseTimes(getCommandLineInput())
	sortedOrder = bst.SortedOrder()
	Percentile99()
	Percentile95()
	Percentile90()
	Percentile1()
}
