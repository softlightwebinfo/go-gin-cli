package internal

import (
	"cli/code/config"
)

type template struct {
}

func NewTemplate() *template {
	return &template{}
}

func (r *template) Env() string {
	return config.Env("postgres", "postgres", "changeme", "localhost", "5432")
}

func (r *template) EnvExample() string {
	return config.Env("", "", "", "", "")
}

func (r *template) Gitignore() string {
	return config.Gitignore()
}

func (r *template) GoMod(nameProject string) string {
	return config.GoMod(nameProject)
}

func (r *template) Main(appName string) string {
	return config.Main(appName)
}
