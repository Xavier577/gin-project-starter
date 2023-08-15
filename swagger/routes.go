package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRoute(r gin.IRouter) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "docs/swagger/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
