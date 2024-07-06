package timeNow

import "time"

func GetTime() string {
	return time.Now().Format(FormatGetTime)
}
