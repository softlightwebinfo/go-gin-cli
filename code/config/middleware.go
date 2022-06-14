package config

import "cli/helpers"

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

func Middleware() string {
	return helpers.Trim(middlewareData)
}
