# config 模块说明

本部分主要负责系统配置的初始化工作

## 文件结构
config/
├── config.md
├── mysql.go  
├── redis.go       
└── securekey.go         


## Mysql数据库初始化(InitMySQL)

步骤为:
1.创建一个数据库连接；
2.自动创建 `student_db` 数据库（如果不存在）；
3.重新将数据库连接并连接到 `student_db` 上;
4.创建表单(内容包括id, name, tel, study);
5.连接对象保存在 `config.DB` 中供全局调用；

## Redis初始化(InitRedis)

步骤为:
1.建立 Redis 客户端连接；
2.Redis 使用 `config.RDB` 和上下文 `config.Ctx`；
3.默认连接本地 Redis 服务（localhost:6379）；
4.用于缓存学生信息，加速系统读取；

## 初始化查询密钥(InitSecurekey)

加载并读取密钥

## 注释

为了便于使用，我没有设置mysql与redis的密码，使用时应修改
- "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
- "root:@tcp(127.0.0.1:3306)/student_db"
-  Addr:     "localhost:6379", // Redis 地址和端口
-  Password: "",               // 没有设置密码
-  DB:       0,                // 使用默认数据库 0