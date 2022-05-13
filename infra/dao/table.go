package dao

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Product product
type Product struct {
	ID      int64 `gorm:"primary_key"`
	Name    string
	BrandID int64
	BaseModel
}

// TableName テーブル名定義
func (Product) TableName() string {
	return "product"
}

// Brand brand
type Brand struct {
	ID   int64 `gorm:"primary_key"`
	Name string
	BaseModel
}

// TableName テーブル名定義
func (Brand) TableName() string {
	return "brand"
}
