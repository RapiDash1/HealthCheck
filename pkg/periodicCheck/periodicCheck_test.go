package periodicCheck

import "testing"

// Test
func TestPeriodic(t *testing.T) {
	url := "https://www.youtube.com/"
	duration := 0.5
	logLoc := "../../logs/"
	periodicCheck(url, duration, logLoc)
}
