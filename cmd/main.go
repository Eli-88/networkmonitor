package main

import (
	"networkmonitor/cmd/app"
	"networkmonitor/cmd/config"
	"networkmonitor/cmd/handler"
	db "networkmonitor/core/db/kv"
	"networkmonitor/core/logger"
	"networkmonitor/core/timer"
	"os"
)

func main() {
	logger.SetLogLevel(logger.TRACE)

	if len(os.Args) < 2 {
		logger.Fatal("need at least 1 arguement for config path. sample usage: go run cmd/main.go <config path>")
	}

	config, err := config.MakeConfig(os.Args[1])
	if err != nil {
		logger.Fatal(err)
	}

	db, err := db.MakeBadgerDb(config.DbPath())
	if err != nil {
		logger.Fatal(err)
	}
	builder := app.MakeBuilder(config.ServerIpAddr(), db, timer.MakeTimer(), config)
	pingEngine := builder.BuildPingEngine()
	rankEngine := builder.BuildRankEngine()
	handlerBuilder := handler.MakeHandlerBuilder(pingEngine, rankEngine)

	app := app.MakeApp(
		builder.BuildHttpServer(),
		pingEngine,
		rankEngine,
		handlerBuilder)

	app.Run()
}
