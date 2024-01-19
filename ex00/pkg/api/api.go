// package api
package api

import (
	"github.com/gin-gonic/gin"
	"omikuji/pkg/omikuji"
)

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// おみくじの結果を返す
		result := omikuji.RandomOmikuji()
		c.JSON(200, result)
	})
	return r
}