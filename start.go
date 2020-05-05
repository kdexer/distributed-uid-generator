package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kdexer/distributed-uid-generator/config"
	"github.com/kdexer/distributed-uid-generator/generator"
	"github.com/kdexer/distributed-uid-generator/routers"
	"net/http"
)

/**
 * 启动函数
 */
func main() {
	yamlConfig := config.ReadConfig(config.CONFIG_FILE_PATH)
	Config := yamlConfig.Config
	// todo 处理worker的问题
	dg := generator.New(1,
		Config.Date.EpochDate,
		Config.Bits.Time,
		Config.Bits.Worker,
		Config.Bits.Sequences)
	generatorRouter := routers.NewGeneratorRouter(dg)
	genHandler := generatorRouter.GetHandlerFunction()
	router := httprouter.New()
	router.GET(Config.Server.Path.Generator, genHandler)
	// todo 处理的不好，以后有时间处理
	startError := http.ListenAndServe(":"+Config.Server.Port, router)
	if nil != startError {
		panic("server start up error")
	}
}
