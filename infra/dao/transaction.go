package dao

import (
	"fmt"

	"github.com/haradayoshitsugucz/purple-server/domain/repository"
	"github.com/haradayoshitsugucz/purple-server/logger"
	"gorm.io/gorm"
)

type TransactionDao struct {
	writer *gorm.DB
}

func NewTransactionRepository(cluster *DBCluster) repository.TransactionRepository {
	return &TransactionDao{writer: cluster.writer}
}

func (db *TransactionDao) Begin() (*gorm.DB, error) {

	logger.GetLogger().Debug("Begin Transaction")

	tx := db.writer.Begin()
	if err := tx.Error; err != nil {
		logger.GetLogger().Error(fmt.Sprintf("error on begin transaction: %+v", err))
		panic(err)
	}
	return tx, tx.Error
}

func (db *TransactionDao) Commit(tx *gorm.DB) error {

	logger.GetLogger().Debug("Commit Transaction")
	tx.Commit()
	if err := tx.Error; err != nil {
		logger.GetLogger().Error(fmt.Sprintf("error on commit transaction: %+v", err))
		panic(err)
	}
	return tx.Error
}

func (db *TransactionDao) Rollback(tx *gorm.DB) error {

	logger.GetLogger().Warn("Rollback Transaction")

	tx.Rollback()
	if err := tx.Error; err != nil {
		logger.GetLogger().Error(fmt.Sprintf("error on rollback transaction: %+v", err))
		panic(err)
	}
	return tx.Error
}
