package config

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	CONFIG_FILE_PATH string = "./resources/config"
)

/**
 * 配置信息结构体
 */
type config struct {
	epochDate string
	timebit int32
	workbit int32
	sequencesbit int32
}

func (c config) Sequencesbit() int32 {
	return c.sequencesbit
}

func (c config) Workbit() int32 {
	return c.workbit
}

func (c config) Timebit() int32 {
	return c.timebit
}

func (c config) EpochDate() string {
	return c.epochDate
}



/**
 * 创建配置对象
 */
func New() *config {
	config := new(config)
	configMap := getConfigMap()
	//todo 利用反射处理，单独写一个函数
	timebit, e := strconv.ParseUint(configMap["timebit"], 10, 8)
	if (e != nil) {
		panic("config timebit key parse error")
	}

	workbit, e := strconv.ParseUint(configMap["workbit"], 10, 8)
	if (e != nil) {
		panic("config workbit key parse error")
	}
	sequencesbit, e := strconv.ParseUint(configMap["sequencesbit"], 10, 8)
	if (e != nil) {
		panic("config sequencesbit key parse error")
	}

	config.epochDate = configMap["epochDate"]
	config.timebit = int32(timebit)
	config.workbit = int32(workbit)
	config.sequencesbit = int32(sequencesbit)

	return config
}

/**
 * 获取文件读取对象
 */
func getOsReader() *bufio.Reader {
	file, e := os.Open(CONFIG_FILE_PATH)
	if (nil != e || nil == file) {
		//todo 添加日志
		panic("读取配置获取OSReader失败:" +e.Error())
	}
	reader := bufio.NewReader(file)
	return reader
}

/**
 * 解析文件内容到map容器中
 */
func getConfigMap()  map[string]string{
	reader := getOsReader()
	configMap := make(map[string]string)
	for true {
		line, _, err := reader.ReadLine()
		if (nil != err && io.EOF != err) {
			panic("配置文件读取失败")
		}
		if (nil != err && io.EOF == err) {
			break
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

