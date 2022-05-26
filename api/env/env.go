package env

import (
	"fmt"
	"log"

	"github.com/haradayoshitsugucz/purple-server/config"
)

// NewConfig config初期化
func NewConfig(env string) config.Config {
	switch env {
	case config.Local:
		return NewLocal()
	case config.Test:
		return NewTest()
	default:
		log.Panic(fmt.Sprintf("env引数が不正なため、起動に失敗しました en%+v\n", env))
	}
	return NewLocal()
}
