package helper

import (
	"strconv"
	"time"
)

func FormatTime(t *time.Time) string {
	if t != nil {
		return t.String()
	}
	return ""
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func StringToFloat(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

func StringToBool(s string) bool {
	i, _ := strconv.ParseBool(s)
	return i
}

func StringToTime(s string) time.Time {
	i, _ := time.Parse(time.RFC3339, s)
	return i
}
