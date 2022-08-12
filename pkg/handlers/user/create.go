package user

import (
	"fmt"
	"gin-project-starter/pkg/dtos"
	"gin-project-starter/pkg/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var newUser dtos.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		statusCode := http.StatusBadRequest
		errMsg := err.Error()
		if err.Error() == "EOF" {
			errMsg = "Invalid Request Body"
		}
		ctx.JSON(statusCode, map[string]any{"message": errMsg})
		return
	}

	userExistsWithEmail := user.AlreadyExistsWithEmail(newUser.Email)

	if userExistsWithEmail {
		ctx.JSON(http.StatusBadGateway, map[string]any{"message": fmt.Sprintln("user already Exists with", newUser.Email)})
		return
	}

	err := user.Create(&newUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{"message": "user successfully created"})

}
