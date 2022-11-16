package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

//InitCfg init redis db
func InitCfg() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // 密码,viper取不到的话默认为空
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func Close() {
	_ = rdb.Close()
}
