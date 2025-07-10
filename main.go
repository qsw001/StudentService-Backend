package main

import (
	"log"
	"student-service/config"
	"student-service/routes"
)

//以下的config与router为自己定义的包
func main(){
	//初始数据库
	config.InitMySQL()  
	config.InitRedis()
	defer config.DB.Close()

	//建立一个路由
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil{
		log.Fatal("服务器启动失败", err)
	}

}