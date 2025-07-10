# routes(路由)模块说明

routes 模块负责初始化并配置 Gin 框架的路由映射关系。

## 文件结构

routes/
└── routes.go

## 路由说明

- /students	        GET	        查询所有学生	    controller.ListStudents
- /students	        POST	    创建新学生	        controller.CreateStudent
- /students/:id	    GET	        查询某个学生	    controller.GetStudent
- /students/:id	    PUT	        更新学生信息	    controller.UpdateStudent
- /students/:id	    DELETE      删除学生信息	    controller.DeleteStudent

## 注释

本项目目前只添加了一个路由组，后续可以继续扩展