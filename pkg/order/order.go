package order

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        string
	Task      string
	Details   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Insert(db *gorm.DB, order *Order) {
	order.ID = uuid.NewString()
	err := db.Create(order).Error
	if err != nil {
		panic(err)
	}
}
