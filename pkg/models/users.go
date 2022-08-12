package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	PK       uint   `gorm:"primaryKey"`
	UserID   string `gorm:"type:uuid;unique;"`
	Email    string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID, uuidGenErr := uuid.NewV4()

	if uuidGenErr != nil {
		return uuidGenErr
	}

	u.UserID = newUUID.String()

	return
}
