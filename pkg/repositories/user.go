package repositories

import (
	"gin-project-starter/pkg/dtos"
	"gin-project-starter/pkg/models"
	"log"

	"gorm.io/gorm"
)

type IUser interface {
	FindById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(u *dtos.User) error
	Exists(query interface{}, args ...interface{}) bool
}

type user struct {
	DbClient *gorm.DB
}

func (u *user) Create(newUser *dtos.User) error {

	user := models.User{Email: newUser.Email,
		Password: newUser.Password}

	err := u.DbClient.Create(&user).Error

	return err

}

func (u *user) FindById(id string) (*models.User, error) {
	var user models.User
	err := u.DbClient.Take(&user, models.User{UserID: id}).Error
	return &user, err
}

func (u *user) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DbClient.Take(&user, models.User{Email: email}).Error
	return &user, err
}

// queries the database to see if a user with the query conditon exists (note this first line)
//
// Examples:
//
//	userExistsWithId := userRepo.Exists("id = ?", id)
//	userExistsWithUsernameOrEmail := userRepo.Exists("username = ? or email = ?", username, email)
//
//	if userExistsWithId {
// 		// do something
//	}
//	if userExistsWithUsernameOrEmail {
// 		// do something
//	}
//
func (u *user) Exists(query interface{}, args ...interface{}) bool {
	var exists bool
	err := u.DbClient.Model(&models.User{}).Select("count(*) > 0").Where(query, args...).Find(&exists).Error

	if err != nil {
		log.Fatal(err)
	}

	return exists

}
