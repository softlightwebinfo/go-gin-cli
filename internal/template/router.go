package internal

import "cli/code/config"

func (r *template) Router() string {
	return config.Router()
}
