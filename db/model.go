package db

import (
	"time"

	"gorm.io/gorm"
)

// Order
type Order struct {
	gorm.Model
	ID           string    `gorm:"primaryKey;column:id;type:varchar(36);not null"`
	UserID       string    `gorm:"column:user_id;type:varchar(36)"`
	ProductCount int       `gorm:"column:product_count;type:int(11)"`
	Price        float64   `gorm:"column:price;type:decimal(10,0)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime"`
}

// User
type User struct {
	gorm.Model
	ID         string    `gorm:"primaryKey;column:id;type:varchar(36);not null"`
	UserName   string    `gorm:"column:user_name;type:varchar(12)"`
	Age        int8      `gorm:"column:age;type:tinyint(3)"`
	Phone      string    `gorm:"column:phone;type:varchar(11)"`
	Province   string    `gorm:"column:province;type:varchar(10)"`
	City       string    `gorm:"column:city;type:varchar(10)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
}
