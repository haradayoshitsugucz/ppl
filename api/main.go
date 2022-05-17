package main

import (
	"os"
	"time"

	"github.com/haradayoshitsugucz/purple-server/api/env"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/router"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/urfave/cli/v2"
)

func main() {

	application := &cli.App{
		Name:    "purple-api",
		Usage:   "purple api service",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "env", Value: "local", Usage: "環境を指定して下さい"},
			&cli.StringFlag{Name: "log_dir", Value: "", Usage: "ログディレクトリ名を指定して下さい"},
			&cli.StringFlag{Name: "log_file", Value: "", Usage: "ログファイル名を指定して下さい"},
		},
		Action: func(context *cli.Context) error {
			utc := time.FixedZone("UTC", 0)
			time.Local = utc
			envArg := context.String("env")
			logDirArg := context.String("log_dir")
			logFileNameArg := context.String("log_file")
			log := config.EmptyLog()
			if len(logFileNameArg) > 0 {
				log = &config.LogArgs{
					Dir:      logDirArg,
					FileName: logFileNameArg,
				}
			}
			conf := env.NewConfig(envArg)
			return router.Run(conf, log)
		},
	}

	err := application.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
