package model

import (
	"fmt"
)

type Product struct {
	ID      int64
	Name    string
	BrandID int64
	empty   bool
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
