package config

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB


//创建数据库，创建数据库的表，与go连接
func InitMySQL(){
	dsn1 := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"

	//连接
	db, err := sql.Open("mysql", dsn1)//初始化数据库连接对象
	if err != nil{
		log.Fatal("数据库连接失败：",err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("数据库不可用：",err)
	}

	log.Println("数据库连接成功！")

	//创建数据库
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS student_db DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
    if err != nil {
        log.Fatal("创建数据库失败：", err)
    }

	log.Println("数据库创建成功")

	//重新连接数据库
	dsn := "root:@tcp(127.0.0.1:3306)/student_db"

	db, err = sql.Open("mysql", dsn)
	if err != nil{
		log.Fatal("数据重新库连接失败：",err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()
	if err != nil{
		log.Fatal("数据库不可用：",err)
	}

	log.Println("数据库重新连接成功！")

	//创建表
	  createTableSQL := `
    CREATE TABLE IF NOT EXISTS students (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        tel VARCHAR(20),
        study VARCHAR(100)
    );`

	 _, err = db.Exec(createTableSQL)
    if err != nil {
        panic(err)
    }

	log.Println("建表成功")

	DB = db
}
	
