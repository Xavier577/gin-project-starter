package user

import (
	"gin-project-starter/pkg/dtos"
	"gin-project-starter/pkg/repositories"
)

func Create(newUser *dtos.User) error {
	err := repositories.User.Create(newUser)

	return err
}
