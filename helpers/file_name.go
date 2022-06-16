package helpers

import (
	"regexp"
	"strings"
)

func FileName(name string) string {
	var filenames []string
	r, _ := regexp.Compile("(^[a-z]+[a-z])|([A-Z]+[a-z]*)")

	for _, item := range r.FindAllStringSubmatch(name, -1) {
		name := item[0]
		filenames = append(filenames, strings.ToLower(name))
	}

	return strings.Join(filenames, "_")
}
