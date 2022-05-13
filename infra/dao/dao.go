package dao

import (
	"fmt"

	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
	"github.com/haradayoshitsugucz/purple-server/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

const duplicateEntryErrorCode = 1062

type DBCluster struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewDBCluster(conf config.Config, secretsRepo repository.SecretsRepository) *DBCluster {

	var w *gorm.DB
	if len(conf.DBWriterSetting().SecretName) > 0 {
		writerCredentials, err := secretsRepo.GetDBCredentials(conf.DBWriterSetting().SecretName)
		if err != nil {
			panic(err)
		}
		w = newDB(conf.DBWriterSetting().WithCredentials(writerCredentials), conf.LoggerSetting())
	} else {
		w = newDB(conf.DBWriterSetting(), conf.LoggerSetting())
	}

	var r *gorm.DB
	if len(conf.DBReaderSetting().SecretName) > 0 {
		readerCredentials, err := secretsRepo.GetDBCredentials(conf.DBReaderSetting().SecretName)
		if err != nil {
			panic(err)
		}
		r = newDB(conf.DBReaderSetting().WithCredentials(readerCredentials), conf.LoggerSetting())
	} else {
		r = newDB(conf.DBReaderSetting(), conf.LoggerSetting())
	}

	return &DBCluster{
		writer: w,
		reader: r,
	}
}

func newDB(setting *config.DBSetting, loggerSetting *config.LoggerSetting) *gorm.DB {

	// dbSetting
	dbConnectionString := fmt.Sprintf("%s:%s@%s([%s]:%d)/%s%s",
		setting.User,
		setting.Password,
		setting.Protocol,
		setting.Host,
		setting.Port,
		setting.Name,
		setting.Args,
	)

	l := zapgorm2.New(logger.GetLogger())
	l.SetAsDefault()
	l.IgnoreRecordNotFoundError = loggerSetting.DBLogger.IgnoreRecordNotFoundError

	db, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{
		Logger:      l.LogMode(loggerSetting.DBLogger.Level),
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定します
	sqlDB.SetMaxIdleConns(setting.Params.MaxIdleConns)

	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します（1コネクションプール毎のコネクション）
	sqlDB.SetMaxOpenConns(setting.Params.MaxOpenConns)

	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	sqlDB.SetConnMaxLifetime(setting.Params.ConnMaxLifetime)

	logger.GetLogger().Info(fmt.Sprintf("DB: %s:%d, Params: %+v", setting.Host, setting.Port, setting.Params))

	return db
}

// DuplicateError　duplicate error
type DuplicateError struct {
	message string
}

// Error
func (f *DuplicateError) Error() string {
	return f.message
}
