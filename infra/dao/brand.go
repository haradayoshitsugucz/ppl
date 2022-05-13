package dao

import (
	"errors"

	"github.com/haradayoshitsugucz/purple-server/domain/model"
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

func (db *BrandDao) FindByID(brandID int64) (*model.Brand, error) {

	entity := &Brand{}
	res := db.reader.Where("id = ?", brandID).First(entity)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.EmptyBrand(), nil
	}

	if err := res.Error; err != nil {
		return nil, err
	}

	brand := &model.Brand{
		ID:   entity.ID,
		Name: entity.Name,
	}

	return brand, nil
}

func (db *BrandDao) ListByIDs(brandIDs []int64, limit int) ([]*model.Brand, error) {

	if len(brandIDs) == 0 {
		return []*model.Brand{}, nil
	}

	var entities []*Brand
	res := db.reader.
		Where("id in(?)", brandIDs).
		Find(&entities).
		Limit(limit)

	if err := res.Error; err != nil {
		return nil, err
	}

	brands := make([]*model.Brand, 0, len(entities))
	for _, e := range entities {
		b := &model.Brand{
			ID:   e.ID,
			Name: e.Name,
		}
		brands = append(brands, b)
	}

	return brands, nil
}
