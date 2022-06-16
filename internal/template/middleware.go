package internal

import "cli/code/config"

func (r *template) Middleware() string {
	return config.Middleware()
}

func (r *template) MiddlewareEmpty(name string) string {
	return config.MiddlewareEmpty(name)
}
