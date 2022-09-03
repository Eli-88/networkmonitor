package main

import (
	"networkmonitor/cmd/app"
	"networkmonitor/cmd/handler"
	db "networkmonitor/core/db/kv"
	"networkmonitor/core/logger"
	"networkmonitor/core/timer"
)

func main() {
	logger.SetLogLevel(logger.TRACE)
	db, err := db.MakeBadgerDb("./tmp/badger")
	if err != nil {
		logger.Fatal(err)
	}
	builder := app.MakeBuilder(":5050", db, timer.MakeTimer())
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
