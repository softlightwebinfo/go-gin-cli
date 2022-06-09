package helpers

import "strings"

func Trim(data string) string {
	return strings.Trim(data, "\n")
}
