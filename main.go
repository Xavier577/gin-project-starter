package main

import (
	"fmt"
	_ "github.com/Xavier577/gin-project-starter/docs"
	"github.com/Xavier577/gin-project-starter/pkg/env"
	"github.com/Xavier577/gin-project-starter/swagger"
	"github.com/Xavier577/gin-project-starter/users"
	"github.com/gin-gonic/gin"
	"log"
)

// @title Gin Swagger Api
// @version 1.0
// @description This is a sample swagger doc.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io2

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /v2

func init() {
	ginMode := gin.DebugMode

	AppEnv := env.Get[string]("APP_ENV")

	if AppEnv == env.Production {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)
}

func server() *gin.Engine {
	app := gin.Default()

	api := app.Group("/api")

	swagger.SetupRoute(api)

	userRouter := api.Group(users.UserRoute)

	users.SetRoutes(userRouter)

	return app

}

func main() {

	app := server()

	ADDRESS := fmt.Sprintf(":%s", env.Get[string]("PORT"))

	err := app.Run(ADDRESS)

	if err != nil {
		log.Fatal(err)
	}

}
