package config

import (
	"cli/helpers"
	"fmt"
)

const env = `
DB_DB=%s
DB_USER=%s
DB_PASS=%s
DB_HOST=%s
DB_PORT=%s
ENVIRONMENT="development"
`

func Env(db, user, pass, host, port string) string {
	return helpers.Trim(fmt.Sprintf(env, db, user, pass, host, port))
}
