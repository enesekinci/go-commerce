package helper

import (
	"github.com/gosimple/slug"
	"strconv"
	"strings"
)

func IncrementSlug(name string) string {
	words := strings.Split(name, "-")
	lastWord := words[len(words)-1]

	slugCount := 0

	if lastWord != "" {
		count, err := strconv.Atoi(lastWord)
		if err != nil {
			count = 0
		}

		slugCount = count
	}

	return name + "-" + strconv.Itoa(slugCount+1)
}

func NewSlug(name string) string {
	return slug.Make(name)
}
