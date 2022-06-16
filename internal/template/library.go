package internal

import "cli/code/config"

func (r *template) Library() string {
	return config.Library()
}

func (r *template) LibraryNew(libraryName string) string {
	return config.LibraryNew(libraryName)
}
