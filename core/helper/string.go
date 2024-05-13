package helper

import (
	"github.com/gosimple/slug"
	"strconv"
	"strings"
)

func NewSlug(name string, table string, column string) string {
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

	seo := slug.Make(name) + strconv.Itoa(slugCount)

	if IsExistInDB(table, column, seo) {
		seo = name + "-" + strconv.Itoa(slugCount+1)
		return NewSlug(seo, table, column)
	}

	return seo
}
