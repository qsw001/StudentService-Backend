package main

import (
	"log"
	"student-service/config"
	"student-service/routes"
	"student-service/models"
	"fmt"
	"time"
	 "net/http"
    _ "net/http/pprof" 
)

//以下的config与router为自己定义的包
func main(){
	go func() {
        fmt.Println("pprof listening on :6060")
        http.ListenAndServe("localhost:6060", nil)
    }()


	//初始数据库
	config.InitMySQL()  
	config.InitRedis()
	config.InitSecurekey()

	testMySQLConnectionPool()
	
	defer config.Sqldb2.Close()

	err := config.DB.AutoMigrate(&models.Student{})

	if err != nil {
		log.Fatal("自动建表失败：", err)
	}
	log.Println("自动建表成功")

	//建立一个路由
	r := routes.SetupRouter()
	err = r.Run(":8080")
	if err != nil{
		log.Fatal("服务器启动失败", err)
	}

}


func testMySQLConnectionPool() {
    db := config.DB

    for i := 0; i < 20; i++ {
        go func(i int) {
            var s models.Student
            db.First(&s, 7)

            sqlDB, _ := db.DB()
            stats := sqlDB.Stats()
            fmt.Printf("Goroutine %d: OpenConns = %d, InUse = %d, Idle = %d\n",
                i, stats.OpenConnections, stats.InUse, stats.Idle)
        }(i)
    }

    time.Sleep(5 * time.Second) // 等待所有 goroutine 执行完
}