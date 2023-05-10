package order

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int
	UUID      string
	Task      string
	Details   string
	UserUUID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Insert(db *gorm.DB, order *Order) (*Order, error) {
	order.UUID = uuid.NewString()
	err := db.Create(order).Error
	return order, err
}

func OrderById(db *gorm.DB, order_id string) (Order, error) {
	var order Order
	err := db.Where("id = ?", order_id).First(&order).Error
	return order, err
}

func OrdersByUser(db *gorm.DB, user_id string) ([]Order, error) {
	var orders []Order
	err := db.Where("user_id = ?", user_id).Find(&orders).Error
	return orders, err
}
