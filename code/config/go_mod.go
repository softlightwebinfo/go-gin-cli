package config

import (
	"cli/helpers"
	"fmt"
)

const mod = `
module %s

go 1.18
`

func GoMod(module string) string {
	return helpers.Trim(fmt.Sprintf(mod, module))
}
