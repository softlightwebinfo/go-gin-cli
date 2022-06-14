package internal

import "cli/code/config"

func (r *template) Config() string {
	return config.Config()
}

func (r *template) ConfigDB() string {
	return config.ConfigDB()
}

func (r *template) ConfigEmail() string {
	return config.ConfigEmail()
}
