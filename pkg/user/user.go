package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Insert(db *gorm.DB, user *User) *User {
	user.ID = uuid.NewString()
	err := db.Create(user).Error
	if err != nil {
		panic(err)
	}
	return user
}

func UserById(db *gorm.DB, id string) User {
	var user User
	db.Model(User{ID: id}).First(&user)
	return user
}
