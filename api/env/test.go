package env

import (
	"time"

	"github.com/haradayoshitsugucz/purple-server/config"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
)

// test
type test struct {
}

func NewTest() *test {
	return &test{}
}

func (e *test) Description() string {
	return "Starting API application by TEST configuration."
}

func (e *test) AppSetting() *config.AppSetting {
	return &config.AppSetting{
		ContextPath: "api",
		Port:        8081,
	}
}

func (e *test) DBWriterSetting() *config.DBSetting {
	return &config.DBSetting{
		User:     "root",
		Password: "",
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     3307,
		Name:     "test_purple",
		Args:     "?parseTime=true&loc=UTC&rejectReadOnly=true",
		Params: &config.DBParams{
			MaxIdleConns:    10,
			MaxOpenConns:    50,
			ConnMaxLifetime: time.Hour,
		},
	}
}

func (e *test) DBReaderSetting() *config.DBSetting {
	return &config.DBSetting{
		User:     "root",
		Password: "",
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     3307,
		Name:     "test_purple",
		Args:     "?parseTime=true&loc=UTC&rejectReadOnly=true",
		Params: &config.DBParams{
			MaxIdleConns:    10,
			MaxOpenConns:    50,
			ConnMaxLifetime: time.Hour,
		},
	}
}

func (e *test) LoggerSetting() *config.LoggerSetting {
	return &config.LoggerSetting{
		Level:         zapcore.DebugLevel,
		Encoding:      "console",
		LogDir:        "/var/log/purple",
		FileName:      "application.log",
		RequestOutput: true,
		DBLogger: &config.DBLoggerSetting{
			Level:                     gormLogger.Info,
			IgnoreRecordNotFoundError: true,
		},
		Rotate: &config.LogRotateSetting{},
	}
}

func (e *test) SecretsManagerSetting() *config.SecretsManagerSetting {
	return &config.SecretsManagerSetting{
		Region:  "ap-northeast-1",
		Profile: "purple-test",
	}
}
