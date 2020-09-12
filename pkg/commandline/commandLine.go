package commandline

import "flag"

// GetUrlInput
func GetUrlInput() string {
	url := flag.String("url", "https://docs.google.com/", "Url to perform health check")
	flag.Parse()
	return *url
}

// GetDurationInput
func GetDurationInput() float64 {
	duration := flag.Float64("duration", 10, "Health check duration")
	flag.Parse()
	return *duration
}

// GetLogLocInput
func GetLogLocInput() string {
	loc := flag.String("log", "./logs/", "Url to perform health check")
	flag.Parse()
	return *loc
}
