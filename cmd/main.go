package main

import (
	"networkmonitor/cmd/app"
	"networkmonitor/cmd/config"
	"networkmonitor/cmd/handler"
	db "networkmonitor/db/kv"
	"networkmonitor/logger"
	"os"
)

func main() {
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
	builder := app.MakeBuilder(db, config)
	pingEngine := builder.BuildPingEngine()
	rankEngine := builder.BuildRankEngine()
	handlerBuilder := handler.MakeHandlerBuilder(pingEngine, rankEngine)

	app := app.MakeApp(
		builder.BuildServer(),
		pingEngine,
		rankEngine,
		handlerBuilder)

	app.Run()
}
