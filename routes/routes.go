package routes

import (
	"student-service/controller"
	"student-service/middleware"
	"github.com/gin-gonic/gin"
)

/*
1. ListStudents所有学生信息的拉取；
2. CreateStudent 创建一个学生的信息；
3. GetStudent 拉取一个学生的信息；
4. UpdateStudent 更新一个学生的信息；
5. DeleteStudent 删除一个学生的信息；
*/

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//注册cors，cors注册要在登陆的前面，防止登陆时就出现跨域请求失败
	r.Use(middleware.CORSMiddleware())

	//注册日志，注册日志也要放在前面，可以记录所有的信息
	r.Use(middleware.AccessLog())

	//加入JWT登陆路由
	r.POST("/login",controller.Login)
	

	// 创建一个以 /students 开头的路由组
	studentGroup := r.Group("/students",middleware.JWTAuthMiddleware())
	{
		studentGroup.GET("", controller.ListStudents)
		studentGroup.POST("", controller.CreateStudent)
		studentGroup.GET("/:id", controller.GetStudent)
		studentGroup.PUT("/:id", controller.UpdateStudent)
		studentGroup.DELETE("/:id", controller.DeleteStudent)
	}

	return r
}
