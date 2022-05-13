package env

import (
	"fmt"

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
		panic(fmt.Errorf("env引数が不正なため、起動に失敗しました env: %v\n", env))
	}
}
