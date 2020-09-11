package periodicCheck

import (
	"fmt"
	"healthCheck/internal/commandLine"
	"healthCheck/pkg/request"
	"log"
	"os"
	"time"
)

// Perform Check
func performCheck(url string) time.Duration {
	return request.ResponseTime(url)
}

// Log Name
func logName(logLoc string) string {
	now := time.Now()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()
	return logLoc + "file_" + fmt.Sprint(hour) + "h" + fmt.Sprint(min) + "m" + fmt.Sprint(sec) + ".log"
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
		_, err := log.WriteString(fmt.Sprintf("%s - Response time -> %s\n", fmt.Sprint(time.Now()), performCheck(url)))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Minute)
	}
}

// Check
func Check() {
	url := commandLine.GetUrlInput()
	duration := commandLine.GetDurationInput()
	logLoc := commandLine.GetLogLocInput()
	periodicCheck(url, duration, logLoc)
}
