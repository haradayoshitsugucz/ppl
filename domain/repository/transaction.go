package repository

import "gorm.io/gorm"

type TransactionRepository interface {
	Begin() (*gorm.DB, error)
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}
