package config

import "cli/helpers"

const configDBData string = `
package config

import(
	"database/sql"
)

var DB *sql.DB
`

func ConfigDB() string {
	return helpers.Trim(configDBData)
}
