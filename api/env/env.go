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

// NewConfig2 config初期化
func NewConfig2(env string) config.Config {
	var conf config.Config
	switch env {
	case config.Local:
		conf = NewLocal()
	case config.Test:
		conf = NewTest()
	default:
		panic(fmt.Errorf("env引数が不正なため、起動に失敗しました env: %v\n", env))
	}
	return conf
}

// NewConfig3 config初期化
func NewConfig3(env string) config.Config {
	var conf config.Config
	switch env {
	case config.Local:
		conf = NewLocal()
	case config.Test:
		conf = NewTest()
	default:
		panic(fmt.Errorf("env引数が不正なため、起動に失敗しました env: %v\n", env))
	}
	return conf
}