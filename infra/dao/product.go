package dao

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/haradayoshitsugucz/purple-server/domain/model"
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

func (db *ProductDao) FindByID(productID int64) (*model.Product, error) {

	entity := &Product{}
	res := db.reader.Where("id = ?", productID).First(entity)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.EmptyProduct(), nil
	}

	if err := res.Error; err != nil {
		return nil, err
	}

	product := &model.Product{
		ID:      entity.ID,
		Name:    entity.Name,
		BrandID: entity.BrandID,
	}

	return product, nil
}

func (db *ProductDao) ListByName(name string, offset, limit int) ([]*model.Product, error) {

	var entities []*Product

	res := db.reader.
		Where("name like ?", fmt.Sprintf("%%%s%%", name)).
		Offset(offset).Limit(limit).
		Order("id DESC").
		Find(&entities)

	if err := res.Error; err != nil {
		return nil, err
	}

	products := make([]*model.Product, 0, len(entities))
	for _, e := range entities {
		p := &model.Product{
			ID:      e.ID,
			Name:    e.Name,
			BrandID: e.BrandID,
		}
		products = append(products, p)
	}

	return products, nil
}

func (db *ProductDao) CountByName(name string) (count int64, err error) {

	res := db.reader.Model(&Product{}).
		Where("name like ?", fmt.Sprintf("%%%s%%", name)).
		Count(&count)

	if err := res.Error; err != nil {
		return 0, err
	}

	return count, res.Error
}

func (db *ProductDao) Insert(product *model.Product, tx *gorm.DB) (productID int64, err error) {

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = db.writer
	}

	entity := &Product{
		Name:    product.Name,
		BrandID: product.BrandID,
	}

	res := conn.Model(&Product{}).Create(entity)
	if err := res.Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == duplicateEntryErrorCode {
				// Error 1062: Duplicate entry
				return 0, &DuplicateError{message: err.Error()}
			}
		}
		return 0, err
	}

	return entity.ID, res.Error
}

// Update update
func (db *ProductDao) Update(product *model.Product, tx *gorm.DB) (rowsAffected int64, err error) {

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

	res := conn.Model(&Product{}).Where("id = ?", product.ID).Updates(updatesMap)

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

	res := conn.Delete(&Product{}, "id = ?", productID)
	if res.RowsAffected == 0 {
		logger.GetLogger().Info("delete product rows_affected is 0")
	}

	return res.RowsAffected, res.Error
}
