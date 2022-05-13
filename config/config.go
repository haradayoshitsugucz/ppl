package config

import (
	"time"

	"github.com/haradayoshitsugucz/purple-server/domain/entity"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
)

const (
	Local = "local"
	Test  = "test"
)

// Config Config
type Config interface {
	Description() string
	AppSetting() *AppSetting
	DBWriterSetting() *DBSetting
	DBReaderSetting() *DBSetting
	LoggerSetting() *LoggerSetting
	SecretsManagerSetting() *SecretsManagerSetting
}

// AppSetting APP設定
type AppSetting struct {
	ContextPath string
	Port        int64
}

// DBSetting DB情報
// SecretNameが指定されている場合、aws secrets manager から User, Password, Host, Port を取得します
type DBSetting struct {
	SecretName string
	User       string
	Password   string
	Protocol   string
	Host       string
	Port       int64
	Name       string
	Args       string
	Params     *DBParams
}

// WithCredentials　secrets manager の値を指定する場合に使用
func (d *DBSetting) WithCredentials(credentials *entity.DBCredentials) *DBSetting {
	d.User = credentials.Username
	d.Password = credentials.Password
	d.Host = credentials.Host
	d.Port = credentials.Port
	return d
}

// DBParams DBパラメータ情報
type DBParams struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// LoggerSetting ログ設定
type LoggerSetting struct {
	Level         zapcore.Level
	Encoding      string
	LogDir        string
	FileName      string
	RequestOutput bool              // リクエストログの出力の有無
	DBLogger      *DBLoggerSetting  // クエリーログの出力の有無
	Rotate        *LogRotateSetting // ログローテートの有無
}

// DBLoggerSetting DBログ設定
type DBLoggerSetting struct {
	Level                     gormLogger.LogLevel // ログを無効にする場合：Silent, ログを有効にする場合：Info/Warn/Error
	IgnoreRecordNotFoundError bool                // [record not found]を出力するかどうか
}

// LogRotateSetting ログローテート設定
type LogRotateSetting struct {
	MaxSize    int // MB
	MaxBackups int
	MaxAge     int // days
	Compress   bool
}

// SecretsManagerSetting SecretsManager設定
type SecretsManagerSetting struct {
	Region  string
	Profile string
}
