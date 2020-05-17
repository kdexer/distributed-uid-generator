package config

import (
	"github.com/go-redis/redis"
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
				Generator string `yaml:"generator"`
			}
		}
		Redis struct{
			Server string `yaml:"server"`
			Password string `yaml:"password"`
			DB int  `yaml:"db"`
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


/*
	从yamlConfig读取配置信息
 */
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

/*
	获取redis客户端
 */
func GetRedisClient(config *YamlConfig) *redis.Client {
	redisConfig := config.Config.Redis
	options := new(redis.Options)
	options.Addr = redisConfig.Server
	options.Password = redisConfig.Password
	options.DB = redisConfig.DB
	client := redis.NewClient(options)
	return client
}