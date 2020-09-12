package periodiccheck

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/RapiDash1/healthCheck/pkg/commandline"
	"github.com/RapiDash1/healthCheck/pkg/request"
)

// Perform Check
func performCheck(url string) int64 {
	return request.ResponseTime(url)
}

// Log Name
func logName(logLoc string) string {
	logName := strings.Split(fmt.Sprint(time.Now()), ".")[0]
	return logLoc + strings.Replace(logName, ":", "-", -1) + ".log"
}

// Periodic Check
func periodicCheck(url string, duration float64, logLoc string) {
	_, err := os.Stat(logLoc)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(logLoc, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	log, err := os.Create(logName(logLoc))
	defer log.Close()
	if err != nil {
		panic(err)
	}
	startTIme := time.Now()
	for {
		if float64(time.Now().Sub(startTIme))/(1e9*60) >= duration {
			return
		}
		_, err := log.WriteString(fmt.Sprintf("%s - Response time -> %dms\n", fmt.Sprint(time.Now()), performCheck(url)))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

// Check health periodically
func Check() {
	url := commandline.GetUrlInput()
	duration := commandline.GetDurationInput()
	logLoc := commandline.GetLogLocInput()
	periodicCheck(url, duration, logLoc)
}
