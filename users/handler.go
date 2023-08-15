package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"id": 1, "email": "johndoe@mail.com"})
}
