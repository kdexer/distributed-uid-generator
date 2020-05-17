package bits

import (
	"github.com/go-redis/redis"
	"github.com/kdexer/distributed-uid-generator/config"
	"strconv"
)

const REDIS_KEY_WORKID = "workId"

func GetWorkId(config *config.YamlConfig,redisCli *redis.Client) int64 {
	workerBits := config.Config.Bits.Worker
	maxWorkerId := -1 ^ (-1 << workerBits)
	workId := redisCli.Get(REDIS_KEY_WORKID)
	if nil == workId || "" == workId.String() {
		result := redisCli.Set(REDIS_KEY_WORKID, 0, 0)
		err := result.Err()
		if nil != err {
			panic("redis设置workid的key失败")
		}
	}

	currentValStr := redisCli.Get(REDIS_KEY_WORKID)
	currentVal, _ := strconv.Atoi(currentValStr.Val())
	if currentVal >= maxWorkerId {
		panic("redis设置workid的key失败,workId超出位数分配范围")
	} else {
		incr := redisCli.IncrBy(REDIS_KEY_WORKID, 1)
		return incr.Val()
	}
}