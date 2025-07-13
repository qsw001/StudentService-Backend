# controller 模块说明

该模块是项目的控制层，负责接收和响应 HTTP 请求。

## 文件结构
controller/
├── auth.go
├── controller.md
└── student_c.go

## 功能介绍

### student_c.go
- ListStudents：获取所有学生数据（GET /students）
- CreateStudent：创建一个新学生记录（POST /students）
- GetStudent：根据 ID 查询单个学生信息（GET /students/:id）
- UpdateStudent：根据 ID 更新学生信息（PUT /students/:id）
- DeleteStudent：根据 ID 删除学生记录（DELETE /students/:id）

### auth.go
- Login: 检验密码是否正确，如果正确返回一个 Token
