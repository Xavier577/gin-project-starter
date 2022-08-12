package routes

import (
	"net/http"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gin-project-starter/docs"
	"gin-project-starter/pkg/handlers/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r gin.IRoutes) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "authorization", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// use ginSwagger middleware to serve the API docs
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "docs/swagger/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/user/create", user.CreateUser)

}
