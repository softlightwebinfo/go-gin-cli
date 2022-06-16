package config

import (
	"cli/helpers"
	"fmt"
)

const middlewareData string = `
package middleware

import(
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine){
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
`

const middlewareEmptyData string = `
package middleware

import (
	"github.com/gin-gonic/gin"
)

func %s () gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
`

func Middleware() string {
	return helpers.Trim(middlewareData)
}

func MiddlewareEmpty(name string) string {
	return helpers.Trim(fmt.Sprintf(middlewareEmptyData, name))
}
