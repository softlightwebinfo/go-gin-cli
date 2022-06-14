package internal

import "cli/code/config"

func (r *template) Middleware() string {
	return config.Middleware()
}
