package commandLine

import "flag"

// Get Url Command Line
func GetUrlInput() string {
	url := flag.String("url", "https://docs.google.com/", "Url to perform health check")
	flag.Parse()
	return *url
}

// Get Duration Command Line Input
func GetDurationInput() float64 {
	duration := flag.Float64("duration", 10, "Health check duration")
	flag.Parse()
	return *duration
}

// Get Log lLoc Command Line Input
func GetLogLocInput() string {
	loc := flag.String("log", "./logs/", "Url to perform health check")
	flag.Parse()
	return *loc
}
