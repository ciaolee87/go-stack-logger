package logReceiver

import (
	"fmt"
	"time"
)

func getTimeString(date time.Time) string {
	return fmt.Sprintf("%02d-%02d-%02dT%02d:%02d:%02dZ", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
}
