package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Sqldb2 *sql.DB

//创建数据库，创建数据库的表，与go连接
func InitMySQL(){
	dsn1 := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"

	//连接
	sqldb, err := sql.Open("mysql", dsn1)//初始化数据库连接对象
	if err != nil{
		log.Fatal("数据库连接失败：",err)
	}

	err = sqldb.Ping()
	if err != nil{
		log.Fatal("数据库不可用：",err)
	}

	log.Println("数据库连接成功！go-sql")

	//创建数据库
	_, err = sqldb.Exec("CREATE DATABASE IF NOT EXISTS student_db DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
    if err != nil {
        log.Fatal("创建数据库失败：", err)
    }

	log.Println("数据库创建成功")

	sqldb.Close()

	//重新连接数据库（使用gorm）
	dsn := "root:@tcp(127.0.0.1:3306)/student_db"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("gorm连接数据库失败：",err)
	}

	sqldb1, err := db.DB()
	if err != nil{
		log.Fatal("获取原生连接池失败")
	}

	//设置连接池参数
	sqldb1.SetMaxOpenConns(50)
	sqldb1.SetMaxIdleConns(10)
	sqldb1.SetConnMaxLifetime(time.Hour)

	Sqldb2 = sqldb1

	log.Println("使用grom连接成功")

	DB = db
}
	
