package repository

import (
	"github.com/haradayoshitsugucz/purple-server/domain/entity"
)

type BrandRepository interface {
	FindByID(brandID int64) (*entity.Brand, error)
	ListByIDs(brandIDs []int64, limit int) ([]*entity.Brand, error)
}
