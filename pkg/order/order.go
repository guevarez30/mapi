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

func OrderById(db *gorm.DB, order_id string) Order {
	var order Order
	db.Model(Order{ID: order_id}).First(&order)
	return order
}

func OrdersByUser(db *gorm.DB) []Order {
	var orders []Order
	db.Model(Order{UserId: "518031f7-bac1-43ba-b5fb-a6045b2e09de"}).Find(&orders)
	return orders
}
