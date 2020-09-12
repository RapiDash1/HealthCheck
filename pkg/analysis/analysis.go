package analysis

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/charts"
)

type respTimes struct {
	keys   []string
	values []float64
}

var responseTimes respTimes

// AnalyzeTodaysLog generates bar graphs for todays logs
func AnalyzeTodaysLog(loc string) {
	folder, err := ioutil.ReadDir(loc)
	if err != nil {
		panic(err)
	}
	nowDate := strings.Split(fmt.Sprint(time.Now()), " ")[0]
	for _, file := range folder {
		extensionlessFileName := strings.Split(file.Name(), ".")[0]
		fileDate := strings.Split(extensionlessFileName, " ")[0]
		if nowDate != fileDate {
			continue
		} else {
			openedFile, err := ioutil.ReadFile(loc + file.Name())
			if err != nil {
				panic(err)
			}
			regTimes, _ := regexp.Compile(" \\d+:\\d+:\\d+")
			loggedTimes := regTimes.FindAllString(string(openedFile), -1)
			regRespTimes, _ := regexp.Compile("\\d+.\\d+ms")
			respTimes := regRespTimes.FindAllString(string(openedFile), -1)

			for index := range loggedTimes {
				responseTimes.keys = append(responseTimes.keys, loggedTimes[index])
				respTime, _ := strconv.ParseFloat(respTimes[index][:len(respTimes[index])-2], 64)
				responseTimes.values = append(responseTimes.values, respTime)
			}
		}
	}
	chartGraph("../../visualization/", nowDate)
}

func chartGraph(loc string, fileName string) {
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
