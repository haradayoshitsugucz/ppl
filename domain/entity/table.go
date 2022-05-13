package entity

import (
	"fmt"
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
	empty bool
}

// TableName テーブル名定義
func (Product) TableName() string {
	return "product"
}

func EmptyProduct() *Product {
	return &Product{empty: true}
}

func (p *Product) Empty() (bool, error) {
	if p.empty {
		return true, fmt.Errorf("product is empty")
	}
	return false, nil
}

// Brand brand
type Brand struct {
	ID   int64 `gorm:"primary_key"`
	Name string
	BaseModel
	empty bool
}

// TableName テーブル名定義
func (Brand) TableName() string {
	return "brand"
}

func EmptyBrand() *Brand {
	return &Brand{empty: true}
}

func (p *Brand) Empty() (bool, error) {
	if p.empty {
		return true, fmt.Errorf("brand is empty")
	}
	return false, nil
}
