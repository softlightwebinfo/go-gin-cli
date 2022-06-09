package internal

import (
	"fmt"
	tmp "text/template"
)

type template struct {
}

func NewTemplate() *template {
	return &template{}
}

func (r *template) getFile(file string) {
	t, err := tmp.ParseFiles(file)
	if err != nil {
		fmt.Println("error:internal:template:getFile", err)
		return
	}
	// err = t.Execute(os.Stdout, std1)
}

func (r *template) Env() string {
	return ""
}
