package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	UUID      string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Insert(db *gorm.DB, user *User) *User {
	user.UUID = uuid.NewString()
	err := db.Create(user).Error
	if err != nil {
		panic(err)
	}
	return user
}

func UserById(db *gorm.DB, UUID string) User {
	var user User
	db.Model(User{UUID: UUID}).First(&user)
	return user
}
