package internal

import (
	"fmt"
	"os"
)

type file struct {
	file string
}

func NewFile() *file {
	return &file{}
}

func (r *file) CreateFile(file string) *file {
	f, err := os.Create(file)
	fmt.Println(file)
	if err != nil {
		fmt.Println("error:internal:file:createFile", err)
		defer f.Close()
	}

	r.file = file
	return r
}

func (r *file) SaveFile(file, data string) *file {
	os.WriteFile(file, []byte(data), os.ModePerm)
	return r
}

func (r *file) AppendTemplate(template string) {
	r.SaveFile(r.file, template)
}
