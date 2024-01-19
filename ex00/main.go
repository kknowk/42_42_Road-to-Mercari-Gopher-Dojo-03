// package main
package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"omikuji/pkg/api"
	"os"
	"omikuji/pkg/cli"
)

func main() {
	// demonize
	gin.SetMode(gin.ReleaseMode)

	// parse args
	port, err := cli.ParseArgs(os.Args)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	// setup server
	r := api.SetupServer()
	r.Run(":" + port)
}
