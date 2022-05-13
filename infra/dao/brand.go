package dao

import (
	"errors"

	"github.com/haradayoshitsugucz/purple-server/domain/entity"
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
	"gorm.io/gorm"
)

type BrandDao struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewBrandRepository(cluster *DBCluster) repository.BrandRepository {
	return &BrandDao{
		writer: cluster.writer,
		reader: cluster.reader,
	}
}

func (db *BrandDao) FindByID(brandID int64) (*entity.Brand, error) {

	brand := &entity.Brand{}
	res := db.reader.Where("id = ?", brandID).First(brand)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return entity.EmptyBrand(), nil
	}

	if err := res.Error; err != nil {
		return nil, err
	}

	return brand, nil
}

func (db *BrandDao) ListByIDs(brandIDs []int64, limit int) ([]*entity.Brand, error) {
	if len(brandIDs) == 0 {
		return []*entity.Brand{}, nil
	}

	var brands []*entity.Brand
	res := db.reader.
		Where("id in(?)", brandIDs).
		Find(&brands).
		Limit(limit)

	if err := res.Error; err != nil {
		return nil, err
	}

	return brands, nil
}
