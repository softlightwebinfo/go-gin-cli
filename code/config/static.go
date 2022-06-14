package config

import "cli/helpers"

const staticData = `
package static

import (
	"github.com/gin-gonic/gin"
)

func Run (r *gin.Engine){
	r.Static("/assets", "./assets")
}
`

func Static() string {
	return helpers.Trim(staticData)
}
