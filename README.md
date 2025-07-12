#  Student-Service 学生信息管理系统（后端）

这是一个基于 **Go（Golang）** 开发的后端项目，使用了 **Gin Web 框架** 搭建 RESTful API，结合 **MySQL 数据库** 实现学生信息的增删查改（CRUD）功能，并使用 **Redis 缓存** 加速数据访问，用JWT鉴权来进行身份验证。

## 项目结构说明

‵``bash
student-service/
├── config/             # 配置模块，MySQL 和 Redis 的初始化
│   ├── mysql.go
│   ├── redis.go
│   ├── securekey.go
│   └── config.md
│
├── controller/         # 控制器模块，处理 HTTP 请求
│   ├── auth.go 
│   ├── student_c.go
│   └── controller.md
│
├── middleware/         # 加入中间层，解析并验证 token
│   ├── auth.go
│   └── middleware.md
│
├── models/             # 模型模块，封装数据库操作逻辑
│   ├── student_m.go
│   └── models.md
│
├── routes/             # 路由注册模块
│   ├── routes.go
│   └── routes.md
│
├── utils/              # 工具函数模块，如 Redis 缓存清理
│   ├── cache.go
│   └── utils.md
│
├── main.go             # 项目入口
├── go.mod              # Go Module 文件
├── go.sum              # Go sum 文件
└── README.md           # 项目说明


## 项目功能概述

-  学生信息的 **增（Create）删（Delete）查（Read）改（Update）**
-  使用 Redis 对查询接口进行 **缓存优化**
-  模块化结构，使用 Gin 实现路由与控制器分离
-  使用 JWT 鉴权来实现身份验证，增加了安全性

## 技术栈

Go	        编程语言
Gin	Web     框架（HTTP服务）
MySQL	    数据库
Redis	    缓存系统

## 使用说明

1. 下载并启动MySQL和Redis
2. 科隆本项目

git clone https://github.com/your-username/student-service.git
cd student-service

3. 安装依赖并运行

go mod tidy
go run main.go
