package repository

import (
	"github.com/haradayoshitsugucz/purple-server/domain/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(productID int64) (*model.Product, error)
	ListByName(name string, offset, limit int) ([]*model.Product, error)
	CountByName(name string) (count int64, err error)
	Insert(product *model.Product, tx *gorm.DB) (productID int64, err error)
	Update(product *model.Product, tx *gorm.DB) (rowsAffected int64, err error)
	DeleteByID(productID int64, tx *gorm.DB) (rowsAffected int64, err error)
}
