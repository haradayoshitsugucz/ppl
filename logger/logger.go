package logger

import (
	"fmt"
	"net/url"
	"os"
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

func InitLogger(conf config.Config, logArgs *config.LogArgs) {

	filePath, err := getFilePath(conf.LoggerSetting(), logArgs)
	if err != nil {
		panic(err)
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

// 設定ファイルと起動引数のどちらも設定した場合は、起動引数の指定が優先される
func getFilePath(setting *config.LoggerSetting, logArgs *config.LogArgs) (string, error) {

	var logDir string
	var filePath string

	// 起動引数の場合
	if ok, _ := logArgs.Empty(); !ok {
		if logArgs.ExistsFile() {
			if logArgs.ExistsDir() {
				logDir = logArgs.Dir
				filePath = fmt.Sprintf("%s/%s", logArgs.Dir, logArgs.FileName)
			} else {
				filePath = logArgs.FileName
			}
		} else {
			return "", fmt.Errorf("[Logger][BootArg] FileName is empty")
		}
	} else {
		// 設定ファイルの場合
		if len(setting.FileName) > 0 {
			if len(setting.LogDir) > 0 {
				logDir = setting.LogDir
				filePath = fmt.Sprintf("%s/%s", setting.LogDir, setting.FileName)
			} else {
				filePath = setting.FileName
			}
		} else {
			if len(setting.LogDir) > 0 {
				return "", fmt.Errorf("[Logger][LoggerSetting] FileName is empty")
			}
		}
	}

	// 事前にログディレクトリを作成する
	if len(logDir) > 0 {
		err := os.MkdirAll(logDir, 0777)
		if err != nil {
			return "", fmt.Errorf("[Logger] Failed to make dir : %s", logDir)
		}
	}

	return filePath, nil
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
		if setting.Rotate != nil && setting.Rotate.MaxSize > 0 {
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
