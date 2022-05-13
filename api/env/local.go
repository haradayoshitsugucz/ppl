package env

import (
	"time"

	"github.com/haradayoshitsugucz/purple-server/config"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
)

// local
type local struct {
}

func NewLocal() *local {
	return &local{}
}

func (e *local) Description() string {
	return "Starting API application by LOCAL configuration."
}

func (e *local) AppSetting() *config.AppSetting {
	return &config.AppSetting{
		ContextPath: "api",
		Port:        8081,
	}
}

func (e *local) DBWriterSetting() *config.DBSetting {
	return &config.DBSetting{
		User:     "root",
		Password: "",
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     3306,
		Name:     "local_purple",
		Args:     "?parseTime=true&loc=UTC&rejectReadOnly=true",
		Params: &config.DBParams{
			MaxIdleConns:    10,
			MaxOpenConns:    50,
			ConnMaxLifetime: time.Hour,
		},
	}
}

func (e *local) DBReaderSetting() *config.DBSetting {
	return &config.DBSetting{
		User:     "root",
		Password: "",
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     3306,
		Name:     "local_purple",
		Args:     "?parseTime=true&loc=UTC&rejectReadOnly=true",
		Params: &config.DBParams{
			MaxIdleConns:    10,
			MaxOpenConns:    50,
			ConnMaxLifetime: time.Hour,
		},
	}
}

func (e *local) LoggerSetting() *config.LoggerSetting {
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

func (e *local) SecretsManagerSetting() *config.SecretsManagerSetting {
	return &config.SecretsManagerSetting{
		Region:  "ap-northeast-1",
		Profile: "purple-local",
	}
}
