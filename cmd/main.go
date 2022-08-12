package main

import (
	"gin-project-starter/pkg/env"
	"gin-project-starter/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	routes.SetRoutes(app)

	app.Run(env.Get("ADDRESS"))
}
