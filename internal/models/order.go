package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusRefunded  OrderStatus = "REFUNDED"
)

type Order struct {
	BaseModel
	UserID      uuid.UUID   `gorm:"type:uuid;not null"`
	OrderNumber string      `gorm:"type:varchar(50);unique;not null"`
	Status      OrderStatus `gorm:"type:varchar(20);not null;default:'PENDING'"`
	TotalAmount float64     `gorm:"type:decimal(12,2);not null"`
	Tax         float64     `gorm:"type:decimal(10,2);not null"`
	ShippingFee float64     `gorm:"type:decimal(10,2);not null"`
	Notes       string      `gorm:"type:text"`
	ShippedAt   *time.Time  `gorm:"type:timestamp"`
	DeliveredAt *time.Time  `gorm:"type:timestamp"`

	User  *User       `gorm:"foreignKey:UserID"`
	Items []OrderItem `gorm:"foreignKey:OrderID"`
}
