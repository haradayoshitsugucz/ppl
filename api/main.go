package main

import (
	"os"
	"time"

	"github.com/haradayoshitsugucz/purple-server/api/env"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/router"
	"github.com/urfave/cli/v2"
)

func main() {

	application := &cli.App{
		Name:    "purple-api",
		Usage:   "purple api service",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "env", Value: "local", Usage: "環境を指定して下さい"},
			&cli.StringFlag{Name: "log", Value: "", Usage: "ファイル名を指定してください"},
		},
		Action: func(context *cli.Context) error {
			utc := time.FixedZone("UTC", 0)
			time.Local = utc
			envArg := context.String("env")
			fileNameArg := context.String("log")
			conf := env.NewConfig(envArg)
			return router.Run(conf, fileNameArg)
		},
	}

	err := application.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
