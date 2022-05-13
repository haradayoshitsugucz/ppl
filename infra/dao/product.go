package dao

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/haradayoshitsugucz/purple-server/domain/entity"
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
	"github.com/haradayoshitsugucz/purple-server/logger"
	"gorm.io/gorm"
)

type ProductDao struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewProductRepository(cluster *DBCluster) repository.ProductRepository {
	return &ProductDao{
		writer: cluster.writer,
		reader: cluster.reader,
	}
}

func (db *ProductDao) FindByID(productID int64) (*entity.Product, error) {

	product := &entity.Product{}
	res := db.reader.Where("id = ?", productID).First(product)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return entity.EmptyProduct(), nil
	}

	if err := res.Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (db *ProductDao) ListByName(name string, offset, limit int) ([]*entity.Product, error) {

	var products []*entity.Product

	res := db.reader.
		Where("name like ?", fmt.Sprintf("%%%s%%", name)).
		Offset(offset).Limit(limit).
		Order("id DESC").
		Find(&products)

	if err := res.Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (db *ProductDao) CountByName(name string) (count int64, err error) {

	res := db.reader.Model(&entity.Product{}).
		Where("name like ?", fmt.Sprintf("%%%s%%", name)).
		Count(&count)

	if err := res.Error; err != nil {
		return 0, err
	}

	return count, res.Error
}

func (db *ProductDao) Insert(product *entity.Product, tx *gorm.DB) (productID int64, err error) {

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = db.writer
	}

	res := conn.Model(&entity.Product{}).Create(product)
	if err := res.Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == duplicateEntryErrorCode {
				// Error 1062: Duplicate entry
				return 0, &DuplicateError{message: err.Error()}
			}
		}
		return 0, err
	}

	return product.ID, res.Error
}

// Update update
func (db *ProductDao) Update(product *entity.Product, tx *gorm.DB) (rowsAffected int64, err error) {

	updatesMap := map[string]interface{}{
		"id":       product.ID,
		"name":     product.Name,
		"brand_id": product.BrandID,
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = db.writer
	}

	res := conn.Model(&entity.Product{}).Where("id = ?", product.ID).Updates(updatesMap)

	if res.RowsAffected == 0 {
		logger.GetLogger().Info("update product rows_affected is 0")
	}

	return res.RowsAffected, res.Error
}

// DeleteByID delete by id
func (db *ProductDao) DeleteByID(productID int64, tx *gorm.DB) (rowsAffected int64, err error) {

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = db.writer
	}

	res := conn.Delete(&entity.Product{}, "id = ?", productID)
	if res.RowsAffected == 0 {
		logger.GetLogger().Info("delete product rows_affected is 0")
	}
	return res.RowsAffected, res.Error
}
