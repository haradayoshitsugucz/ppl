package service

import (
	"log"

	"github.com/haradayoshitsugucz/purple-server/api/domain/dto"
	"github.com/haradayoshitsugucz/purple-server/domain/model"
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
)

type ProductService interface {
	GetProduct(productID int64) (*dto.Product, error)
	GetProductsByName(name string, offset, limit int) (productDTOs []*dto.Product, total int64, err error)
	AddProduct(name string, brandID int64) (productID int64, err error)
	EditProduct(productID int64, name string, brandID int64) (rowsAffected int64, err error)
	DeleteProduct(productID int64) (rowsAffected int64, err error)
}

type ProductServiceImpl struct {
	productRepo     repository.ProductRepository
	brandRepo       repository.BrandRepository
	transactionRepo repository.TransactionRepository
}

func NewProductService(
	productRepo repository.ProductRepository,
	brandRepo repository.BrandRepository,
	transactionRepo repository.TransactionRepository,
) ProductService {
	return &ProductServiceImpl{
		productRepo:     productRepo,
		brandRepo:       brandRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *ProductServiceImpl) GetProduct(productID int64) (*dto.Product, error) {

	// product
	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return nil, err
	}

	if ok, _ := product.Empty(); ok {
		return &dto.Product{Product: model.EmptyProduct(), Brand: model.EmptyBrand()}, nil
	}

	// brand
	brand, err := s.brandRepo.FindByID(product.BrandID)
	if err != nil {
		return nil, err
	}

	if ok, _ := brand.Empty(); ok {
		return &dto.Product{Product: product, Brand: model.EmptyBrand()}, nil
	}

	// DTO
	productDTO := &dto.Product{
		Product: product,
		Brand:   brand,
	}

	return productDTO, nil
}

func (s *ProductServiceImpl) GetProductsByName(name string, offset, limit int) (productDTOs []*dto.Product, total int64, err error) {

	// count
	count, err := s.productRepo.CountByName(name)
	if err != nil {
		return nil, 0, err
	}

	productDTOs = make([]*dto.Product, 0, count)
	if count == 0 {
		return productDTOs, count, nil
	}

	// products
	products, err := s.productRepo.ListByName(name, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	// brandIDs
	brandIDs := make([]int64, 0, count)
	for _, product := range products {
		brandIDs = append(brandIDs, product.BrandID)
	}

	brands, err := s.brandRepo.ListByIDs(brandIDs, len(brandIDs))
	if err != nil {
		return nil, 0, err
	}

	brandMap := make(map[int64]*model.Brand, len(brands))
	for _, b := range brands {
		brandMap[b.ID] = b
	}

	for _, p := range products {
		b := brandMap[p.BrandID]
		if b == nil {
			b = model.EmptyBrand()
		}
		productDTO := &dto.Product{
			Product: p,
			Brand:   b,
		}
		productDTOs = append(productDTOs, productDTO)
	}

	return productDTOs, count, nil
}

func (s *ProductServiceImpl) AddProduct(name string, brandID int64) (productID int64, err error) {

	// transaction
	tx, err := s.transactionRepo.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			s.transactionRepo.Rollback(tx)
		} else {
			s.transactionRepo.Commit(tx)
		}
	}()

	product := &model.Product{
		Name:    name,
		BrandID: brandID,
	}

	productID, err = s.productRepo.Insert(product, tx)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (s *ProductServiceImpl) addProduct2(name string, brandID int64) (productID int64, err error) {

	// transaction
	tx, err := s.transactionRepo.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			if err := s.transactionRepo.Rollback(tx); err != nil {
				log.Print(err)
			}
		} else {
			if err := s.transactionRepo.Commit(tx); err != nil {
				log.Print(err)
			}
		}
	}()

	product := &model.Product{
		Name:    name,
		BrandID: brandID,
	}

	productID, err = s.productRepo.Insert(product, tx)
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (s *ProductServiceImpl) EditProduct(productID int64, name string, brandID int64) (rowsAffected int64, err error) {

	// transaction
	tx, err := s.transactionRepo.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			s.transactionRepo.Rollback(tx)
		} else {
			s.transactionRepo.Commit(tx)
		}
	}()

	product := &model.Product{
		ID:      productID,
		Name:    name,
		BrandID: brandID,
	}
	rowsAffected, err = s.productRepo.Update(product, tx)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *ProductServiceImpl) DeleteProduct(productID int64) (rowsAffected int64, err error) {

	// transaction
	tx, err := s.transactionRepo.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			s.transactionRepo.Rollback(tx)
		} else {
			s.transactionRepo.Commit(tx)
		}
	}()

	rowsAffected, err = s.productRepo.DeleteByID(productID, tx)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
