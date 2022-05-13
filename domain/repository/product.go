package repository

import (
	"github.com/haradayoshitsugucz/purple-server/domain/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(productID int64) (*entity.Product, error)
	ListByName(name string, offset, limit int) ([]*entity.Product, error)
	CountByName(name string) (count int64, err error)
	Insert(product *entity.Product, tx *gorm.DB) (productID int64, err error)
	Update(product *entity.Product, tx *gorm.DB) (rowsAffected int64, err error)
	DeleteByID(productID int64, tx *gorm.DB) (rowsAffected int64, err error)
}
