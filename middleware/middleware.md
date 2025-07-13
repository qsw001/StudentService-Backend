# middleware 模块说明
middleware（中间件）模块是负责解析并验证token令牌的模块

## 文件结构
middleware/
├── auth.go
└── middleware.md

## 功能介绍
负责解析并验证token令牌，成功则将他保存至上下文，失败返回401