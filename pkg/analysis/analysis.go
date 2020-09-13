package analysis

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RapiDash1/healthCheck/pkg/regexhelpers"
	"github.com/go-echarts/go-echarts/charts"
)

var (
	vizFolder string = "../../visualization/"
)

type respTimes struct {
	keys   []string
	values []float64
}

// analyze log files
func analyze(loc string, readAllLogs bool, responseTimes *respTimes) {
	folder, err := ioutil.ReadDir(loc)
	if err != nil {
		panic(err)
	}
	nowDate := strings.Split(fmt.Sprint(time.Now()), " ")[0]
	for _, file := range folder {
		extensionlessFileName := strings.Split(file.Name(), ".")[0]
		fileDate := strings.Split(extensionlessFileName, " ")[0]
		if !readAllLogs && nowDate != fileDate {
			continue
		} else {
			openedFile, err := ioutil.ReadFile(loc + file.Name())
			if err != nil {
				panic(err)
			}
			loggedTimes := regexhelpers.FindAllStringMatches(string(openedFile), " \\d+:\\d+:\\d+")
			respTimes := regexhelpers.FindAllStringMatches(string(openedFile), "\\d+.\\d+ms")

			for index := range loggedTimes {
				responseTimes.keys = append(responseTimes.keys, loggedTimes[index])
				respTime, _ := strconv.ParseFloat(respTimes[index][:len(respTimes[index])-2], 64)
				responseTimes.values = append(responseTimes.values, respTime)
			}
		}
	}
}

// AnalyzeAllLogs generates graph for all logs available
func AnalyzeAllLogs(loc string) {
	var responseTimes respTimes
	analyze(loc, true, &responseTimes)
	chartGraph(vizFolder, "vizAll", &responseTimes)
}

// AnalyzeTodaysLog generates bar graphs for todays logs
func AnalyzeTodaysLog(loc string) {
	var responseTimes respTimes
	analyze(loc, false, &responseTimes)
	nowDate := strings.Split(fmt.Sprint(time.Now()), " ")[0]
	chartGraph(vizFolder, nowDate, &responseTimes)
}

// chartGraph plots a graph at {loc} named as {fileName}.html
func chartGraph(loc string, fileName string, responseTimes *respTimes) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Response time graph"})
	bar.AddXAxis(responseTimes.keys).
		AddYAxis("Response time in ms", responseTimes.values)
	_, err := os.Stat(loc)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(loc, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(loc + fileName + ".html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}
