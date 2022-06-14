package config

import (
	"cli/helpers"
	"fmt"
	"strings"
)

const libraryData = `
package library
`
const libraryNewData = `
package library

type %s struct {}

func New%s() *%s {
	return &%s{}
}
`

func Library() string {
	return helpers.Trim(libraryData)
}

func LibraryNew(libraryName string) string {
	nameNew := strings.Trim(strings.Title(libraryName), " ")
	return helpers.Trim(fmt.Sprintf(libraryNewData, libraryName, nameNew, libraryName, libraryName))
}
