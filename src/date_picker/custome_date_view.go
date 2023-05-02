package date_picker

import "time"

func FormatDataNow() string {
	return time.Now().Format("2006-01-02 03:04")
}
