package main

import (
	"bufio"
	"distributed-uid-generator/generator"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var nextId generator.NextId

func genIdHandler(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uid := nextId.GetNextId()
	fmt.Println(uid)
	fmt.Fprint(rw, uid)
}

// read config info to map
func readConfig(path string) map[string]string {
	file,error := os.Open(path)
	if (nil != error) {
		panic("read config file error")
	}
	configMap := make(map[string]string)
	reader := bufio.NewReader(file)
	for {
		line,_, readErr := reader.ReadLine()
		if nil != readErr {
			if (readErr == io.EOF) {
				break
			}
			panic("read config file error")
		}
		lineStr := string(line)
		index := strings.Index(lineStr, ":")
		// get config key val
		key := lineStr[:index]
		val := lineStr[index+1:]
		configMap[key] = val
	}
	return configMap
}

// start up
func main() {
	config := readConfig("config")

	timebit, e := strconv.ParseUint(config["timebit"], 10, 8)
	if (e != nil) {
		panic("config timebit key parse error")
	}

	workbit, e := strconv.ParseUint(config["workbit"], 10, 8)
	if (e != nil) {
		panic("config workbit key parse error")
	}

	sequencesbit, e := strconv.ParseUint(config["sequencesbit"], 10, 8)
	if (e != nil) {
		panic("config sequencesbit key parse error")
	}

	dg := generator.New(1, config["epochDate"], uint8(timebit), uint8(workbit), uint8(sequencesbit))
	nextId = dg

	router := httprouter.New()
	router.GET("/gen/id", genIdHandler)

	startError := http.ListenAndServe(":8080", router)
	if nil != startError {
		panic("server start up error")
	}
}
