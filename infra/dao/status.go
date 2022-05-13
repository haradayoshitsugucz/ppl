package dao

import (
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
	"gorm.io/gorm"
)

type StatusDao struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewStatusRepository(cluster *DBCluster) repository.StatusRepository {
	return &StatusDao{
		writer: cluster.writer,
		reader: cluster.reader,
	}
}

func (db *StatusDao) Find() error {
	return db.reader.Exec("SELECT 1").Error
}
