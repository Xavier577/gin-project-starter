package users

import (
	"github.com/gin-gonic/gin"
)

const UserRoute = "/users"

func SetRoutes(r gin.IRouter) {
	r.GET("/", getUser)

}
