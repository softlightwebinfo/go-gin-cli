package internal

import (
	"fmt"
	"os"
)

func (r *internal) CreateDir(directory string) *internal {
	if r.directoryRoot != nil {
		directory = fmt.Sprintf("%s/%s", *r.directoryRoot, directory)
		r.directory = directory
	}

	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		fmt.Println("error:internal:create:1", err)
	}
	return r
}

func (r *internal) DirectoryRoot(directory string) *internal {
	r.CreateDir(directory)
	return NewDirectoryRoot(directory)
}

func (r *internal) GetFile(file string) string {
	return fmt.Sprintf("%s/%s", *r.directoryRoot, file)
}

func (r *internal) CreateFile(file string) *file {
	return NewFile().CreateFile(r.GetFile(file))
}
