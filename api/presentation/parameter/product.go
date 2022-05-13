package parameter

import (
	"strconv"
	"strings"

	"github.com/haradayoshitsugucz/purple-server/validation"
)

const (
	defaultOffset = 0
	defaultLimit  = 20
)

// Product
type Product struct {
	ProductID int64 `validate:"min=1"`
}

func (f *Product) Valid() (bool, error) {
	err := validation.Default.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ProductsByName
type ProductsByName struct {
	Name   string `json:"name" validate:"min=1,max=20"`
	Offset int    `json:"offset" validate:"min=0"`
	Limit  int    `json:"limit" validate:"min=1,max=50"`
}

func NewProductsByName(nameValue, offsetValue, limitValue string) *ProductsByName {

	// Name
	name := strings.TrimSpace(nameValue)

	// Offset
	offset, err := strconv.Atoi(offsetValue)
	if err != nil {
		offset = defaultOffset
	}

	// Limit
	limit, err := strconv.Atoi(limitValue)
	if err != nil {
		limit = defaultLimit
	}

	return &ProductsByName{Name: name, Offset: offset, Limit: limit}
}

func (f *ProductsByName) Valid() (bool, error) {
	err := validation.Default.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddProduct
type AddProduct struct {
	Name    string `json:"name" validate:"min=1,max=50"`
	BrandID int64  `json:"brand_id" validate:"min=1"`
}

func (f *AddProduct) Valid() (bool, error) {
	err := validation.Default.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

// EditProduct
type EditProduct struct {
	ProductID int64  `json:"id" validate:"min=1"`
	Name      string `json:"name" validate:"min=1,max=50"`
	BrandID   int64  `json:"brand_id" validate:"min=1"`
}

func (f *EditProduct) Valid() (bool, error) {
	err := validation.Default.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}
