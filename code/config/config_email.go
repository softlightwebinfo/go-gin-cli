package config

import "cli/helpers"

const configEmailData string = `
package config
`

func ConfigEmail() string {
	return helpers.Trim(configEmailData)
}
