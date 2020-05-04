package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/julienschmidt/httprouter"
	"github.com/kdexer/distributed-uid-generator/config"
	"github.com/kdexer/distributed-uid-generator/generator"
	"net/http"
	"os"
)

var nextId generator.NextId

func genIdHandler(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uid := nextId.GetNextId()
	fmt.Println(uid)
	fmt.Fprint(rw, uid)
}

// read config file to YamlConfig
func readConfig(path string) *config.YamlConfig {
	file,readErr := os.Open(path)
	if nil != readErr {
		panic("read config file error")
	}
	yamlConfig := new(config.YamlConfig)
	decoder := yaml.NewDecoder(file)
	decoder.SetStrict(true)
	decoderErr := decoder.Decode(yamlConfig)
	if nil != decoderErr {
		panic("decode yaml config file error")
	}
	return yamlConfig
}

/**
 * 启动函数
 */
func main() {
	yamlConfig := readConfig("resources/config.yaml")
	dg := generator.New(1, yamlConfig.Config.Date.EpochDate,
		yamlConfig.Config.Bits.Time,
		yamlConfig.Config.Bits.Worker,
		yamlConfig.Config.Bits.Sequences)
	nextId = dg
	router := httprouter.New()
	router.GET("/gen/id", genIdHandler)
	startError := http.ListenAndServe(":8080", router)
	if nil != startError {
		panic("server start up error")
	}
}
