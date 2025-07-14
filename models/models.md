# models 模块说明

该模块是项目的model层,用于定义数据结构，并且实现与数据库的交互逻辑 

## 文件结构
```bash
models/
└──student_m.go
```

## 数据结构

ID    int    `json:"id"`
Name  string `json:"name"`
Tel   string `json:"tel"`
Study string `json:"study"`

## 函数介绍
- GetAllStudents() ([]Student, error)
功能：查询所有学生信息。(若启用 Redis，则优先查询缓存，若缓存中无数据再查询数据库，并将数据写入 Redis。)

- GetStudentByID(id int) (Student, error)
功能：根据学生 id 查询单条记录。(若启用 Redis，则优先查询缓存，若缓存中无数据再查询数据库，并将数据写入 Redis。)

- CreateStudent(s Student) (int64, error)
功能：创建一条学生记录。

- UpdateStudent(s Student) error
功能：根据传入 Student.ID 更新数据库中的学生信息。(成功后应清除对应 Redis 缓存。)

- DeleteStudent(id int) error
功能：删除指定 ID 的学生记录。(成功后应清除对应 Redis 缓存。)
