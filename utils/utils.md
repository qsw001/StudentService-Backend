# utils模块说明

`utils` 是本项目中的工具类模块，用于封装通用的辅助函数和功能，使项目更加简洁、易懂。

## 文件结构

utils/
└── cache.go

## 当前功能

- DeleteStudentCache
功能：删除 Redis 中缓存的该学生信息，避免后续读取到旧数据。
(在 UpdateStudent() 和 DeleteStudent() 控制器中使用)