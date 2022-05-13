package repository

import (
	"github.com/haradayoshitsugucz/purple-server/domain/model"
)

type BrandRepository interface {
	FindByID(brandID int64) (*model.Brand, error)
	ListByIDs(brandIDs []int64, limit int) ([]*model.Brand, error)
}
