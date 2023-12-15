package main

import (
	"github.com/fengzhu0601/gotools/cache"
	"github.com/fengzhu0601/gotools/config"
	"github.com/fengzhu0601/gotools/logger"
)

func main() {
	dbConfig := &cache.DBConfig{}
	if err := config.Load("lottery", dbConfig); err != nil {
		panic(err)
	}

	// 启动日志打印
	logger.InitLogger("./log/game_server.log", true)

	logger.Info("DBConfig:", *dbConfig)
}
