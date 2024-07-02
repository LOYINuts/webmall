package cache

import (
	"strconv"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	// 创建redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}

// 初始化
func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	// 读取redis配置
	LoadRedis(file)
	Redis()
}

func LoadRedis(f *ini.File) {
	RedisDb = f.Section("redis").Key("RedisDb").String()
	RedisAddr = f.Section("redis").Key("RedisAddr").String()
	RedisPw = f.Section("redis").Key("RedisPw").String()
	RedisDbName = f.Section("redis").Key("RedisDbName").String()
}
