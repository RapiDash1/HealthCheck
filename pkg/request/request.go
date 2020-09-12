package request

import (
	"fmt"
	"net/http"
	"time"
)

// Response Time
func ResponseTime(url string) int64 {
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
	return diff.Milliseconds()
}
