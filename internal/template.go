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
