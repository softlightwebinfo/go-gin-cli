package config

import "cli/helpers"

const config string = `
package config

var Dev bool = false
`

func Config() string {
	return helpers.Trim(config)
}
