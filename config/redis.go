// config/redis.go
package config

import (
    "context"
    "log"

    "github.com/redis/go-redis/v9"
    "time"
)

var RDB *redis.Client         // Redis 客户端（全局）
var Ctx = context.Background() // 上下文，用于超时控制

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis 地址和端口
        Password: "",               // 没有设置密码
        DB:       0,                // 使用默认数据库 0

        PoolSize:     20,               // 最大连接数（并发高的话调大）
        MinIdleConns: 5,                // 最小空闲连接数
        PoolTimeout:  3 * time.Second, // 等待连接超时时间
    })

    // 测试连接
    _, err := RDB.Ping(Ctx).Result()
    if err != nil {
        log.Fatal("Redis 连接失败：", err)
    }
    

    log.Println("Redis 连接成功！")
}
