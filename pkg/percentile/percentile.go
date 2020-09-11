package percentile

import (
	"fmt"
	"healthCheck/internal/commandLine"
	"healthCheck/pkg/bst"
	"healthCheck/pkg/request"
	"time"
)

var iterations int = 100
var sortedOrder []float64

// Time To Respond
func TimeToRespond(url string, t *chan time.Duration) {
	*t <- request.ResponseTime(url)
}

// Average Time To Respond
func responseTimes(url string) {
	timePerRequest := make(chan time.Duration, iterations)
	for i := 0; i < iterations; i++ {
		go TimeToRespond(url, &timePerRequest)
		time.Sleep(50 * time.Millisecond)
	}
	for i := 0; i < iterations; i++ {
		select {
		case tempTIme := <-timePerRequest:
			bst.AddNode(float64(tempTIme) / 1e9)
		}
	}
}

// Percentile 99
func Percentile99() {
	fmt.Printf("99th percentile: %v\n", sortedOrder[len(sortedOrder)-1])
}

// Percentile 95
func Percentile95() {
	fmt.Printf("95th percentile: %v\n", sortedOrder[len(sortedOrder)-6])
}

// Percentile 90
func Percentile90() {
	fmt.Printf("90th percentile: %v\n", sortedOrder[len(sortedOrder)-11])
}

// Percentile 1
func Percentile1() {
	fmt.Printf("1st percentile: %v\n", sortedOrder[0])
}

// Perform Health Check
func PerformHealthCheck() {
	responseTimes(commandLine.GetUrlInput())
	sortedOrder = bst.SortedOrder()
	Percentile99()
	Percentile95()
	Percentile90()
	Percentile1()
}
