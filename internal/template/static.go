package internal

import (
	"cli/code/config"
)

func (r *template) Static() string {
	return config.Static()
}
