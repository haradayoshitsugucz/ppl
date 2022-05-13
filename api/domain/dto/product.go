package dto

import (
	"github.com/haradayoshitsugucz/purple-server/domain/model"
)

type Product struct {
	Product *model.Product
	Brand   *model.Brand
}
