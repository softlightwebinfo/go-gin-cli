package config

import "cli/helpers"

const makefileData string = `
run:
	nodemon --exec go run main.go --signal SIGTERM
`

func Makefile() string {
	return helpers.Trim(makefileData)
}
