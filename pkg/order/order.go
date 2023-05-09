package order

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Task      string
	Details   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Insert(db *gorm.DB, order *Order) {
	order.ID = uuid.NewString()
	result := db.Create(order)
	if result.Error != nil {
		panic(result.Error)
	}
}

func OrderById(db *gorm.DB, order_id string) (Order, error) {
	var order Order
	result := db.Where("id = ?", order_id).First(&order)
	return order, result.Error

}

func OrdersByUser(db *gorm.DB, user_id string) ([]Order, error) {
	var orders []Order
	result := db.Model(Order{UserId: user_id}).Find(&orders)
	return orders, result.Error
}
