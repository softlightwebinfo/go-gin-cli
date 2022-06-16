package helpers

import "strings"

func FuncName(name string) string {
	return strings.Trim(strings.Title(name), " ")
}
