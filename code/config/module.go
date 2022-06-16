package config

import (
	"cli/helpers"
	"fmt"
)

const moduleRouter = `
package router

import (
	"github.com/gin-gonic/gin"
)

func %s(r *gin.RouterGroup) {

}
`
const moduleController = `
package controller

import "github.com/gin-gonic/gin"

func %s(c *gin.Context) {

}
`
const moduleRepositoryBase = `
package %s

type repository struct{}

func New() *repository {
	return &repository{}
}
`

func ModuleRouter(name string) string {
	return helpers.Trim(fmt.Sprintf(moduleRouter, helpers.FuncName(name)))
}

func ModuleController(name string) string {
	return helpers.Trim(fmt.Sprintf(moduleController, helpers.FuncName(name)))
}

func ModuleBaseRepository(name string) string {
	return helpers.Trim(fmt.Sprintf(moduleRepositoryBase, helpers.FileName(name)))
}
