package config

import "cli/helpers"

const routerData string = `
package router

import(
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine){

}
`

func Router() string {
	return helpers.Trim(routerData)
}
