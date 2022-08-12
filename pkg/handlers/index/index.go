package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"message": "hello"})
}
