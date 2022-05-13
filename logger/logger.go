package logger

import (
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"github.com/haradayoshitsugucz/purple-server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
)

const (
	jsonEncoding    string = "json"
	consoleEncoding string = "console"

	defaultEncoding = consoleEncoding

	// ログファイルの拡張子
	logExt = ".log"
)

func parseEncodingType(in string) string {
	switch in {
	case "json":
		return jsonEncoding
	case "console":
		return consoleEncoding
	default:
		return defaultEncoding
	}
}

func InitLogger(conf config.Config, fileName string) {

	// config指定の場合
	logFile := fmt.Sprintf("%s/%s", conf.LoggerSetting().LogDir, conf.LoggerSetting().FileName)
	if len(fileName) > 0 {
		// 起動引数指定の場合
		logFile = fmt.Sprintf("%s/%s", conf.LoggerSetting().LogDir, fileName)
	}

	// 拡張子が.logの場合はファイル出力をし、それ以外はコンソール出力となる
	ext := filepath.Ext(logFile)
	var filePath string
	if ext == logExt {
		filePath = logFile
	}

	l, err := newZapConfig(
		conf.LoggerSetting(),
		filePath,
	).Build()

	if err != nil {
		panic(err)
	}

	logger = l
	return
}

func GetLogger() *zap.Logger {
	return logger
}

func newZapConfig(setting *config.LoggerSetting, filePath string) zap.Config {

	c := zap.Config{
		Level:       zap.NewAtomicLevelAt(setting.Level),
		Development: false,
		Encoding:    parseEncodingType(setting.Encoding),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig:    newEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if len(filePath) > 0 {
		if setting.Rotate != nil {
			c.OutputPaths = []string{"stdout", fmt.Sprintf("lumberjack:%s", filePath)}
			c.OutputPaths = []string{"stderr", fmt.Sprintf("lumberjack:%s", filePath)}

			setLogWriter(setting.Rotate, filePath)
		} else {
			c.OutputPaths = []string{"stdout", filePath}
			c.ErrorOutputPaths = []string{"stderr", filePath}
		}
	}

	return c
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "eventTime",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    jstTimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02 15:04:05"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func setLogWriter(setting *config.LogRotateSetting, filePath string) {
	ll := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    setting.MaxSize,
		MaxBackups: setting.MaxBackups,
		MaxAge:     setting.MaxAge,
		Compress:   setting.Compress,
	}
	err := zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &ll,
		}, nil
	})

	if err != nil {
		panic(err)
	}
}
