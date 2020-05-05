package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

const (
	CONFIG_FILE_PATH string = "resources/config.yaml"
)

/*
	yaml配置文件的结构体
 */
type YamlConfig struct {
	Config struct {
		Server struct{
			Port string `yaml:"port"`
			Path struct{
				Generator string
			}
		}
		Date struct{
			EpochDate string `yaml:"epochDate"`
		}
		Bits struct{
			Time uint8 `yaml:"time"`
			Worker uint8 `yaml:"worker"`
			Sequences uint8 `yaml:"sequences"`
		}
	}
}


// read config file to YamlConfig
func ReadConfig(path string) *YamlConfig {
	file,readErr := os.Open(path)
	if nil != readErr {
		log.Fatalf("err is %v", readErr)

		panic("read config file error")
	}
	yamlConfig := new(YamlConfig)
	decoder := yaml.NewDecoder(file)
	decoder.SetStrict(true)
	decoderErr := decoder.Decode(yamlConfig)
	if nil != decoderErr {
		panic("decode yaml config file error")
	}
	return yamlConfig
}