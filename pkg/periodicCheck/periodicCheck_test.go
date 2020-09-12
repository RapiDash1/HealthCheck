package periodiccheck

import "testing"

// Test
func TestPeriodic(t *testing.T) {
	url := "https://www.youtube.com/"
	duration := 1.0
	logLoc := "../../logs/"
	periodicCheck(url, duration, logLoc)
}
