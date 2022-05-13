package output

import (
	"github.com/haradayoshitsugucz/purple-server/api/domain/dto"
)

type product struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	BrandName   string `json:"brand_name"`
}

func ToProduct(p *dto.Product) *product {
	return &product{
		ProductID:   p.Product.ID,
		ProductName: p.Product.Name,
		BrandName:   p.Brand.Name,
	}
}

func ToProducts(dtos []*dto.Product) (products []*product) {
	for _, d := range dtos {
		products = append(products, ToProduct(d))
	}
	return products
}
