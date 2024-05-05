package helper

import "time"

func FormatTime(t *time.Time) string {
	if t != nil {
		return t.String()
	}
	return ""
}
