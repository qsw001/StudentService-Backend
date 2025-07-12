package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"student-service/config"
	"time"
)

type Student struct{
	ID    int    `json:"id"`
    Name  string `json:"name"`
    Tel   string `json:"tel"`
    Study string `json:"study"`
}

func GetAllStudents() ([]Student, error){
	//先进行redis查询
	key := "student:all"

	val, err := config.RDB.Get(config.Ctx, key).Result()
	if err == nil{
		var list []Student

		err1 := json.Unmarshal([]byte(val), &list)

		if err1 == nil{
			log.Println("从redis中获取了学生数据")
			return list, err
		}
	}

	//进行数据库查询
	rows, err := config.DB.Query("SELECT id, name, tel, study FROM students")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var list []Student
	for rows.Next() {
		var s Student
		err := rows.Scan(&s.ID, &s.Name, &s.Tel, &s.Study)
		if err != nil{
			return nil, err
		}
		list = append(list, s)
	}

	//写入redis
	data, _ := json.Marshal(list)
	config.RDB.Set(config.Ctx, key, data, 300*time.Second)

	log.Println("从 MySQL 中获取了所有学生数据，并写入 Redis 缓存")

	return list, nil
}

func CreateStudent(s Student) (int64, error){
	result, err := config.DB.Exec("INSERT INTO students(name,tel,study) VALUES(?, ?, ?)",s.Name, s.Tel, s.Study)
	if err != nil{
		return 0,err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0,err
	}

	return id,err
}

func GetStudentByID(id int) (Student, error){
	var s Student

	//先进行redis查询
	key := fmt.Sprintf("student:%d", id)
	val, err := config.RDB.Get(config.Ctx,key).Result()
	if err == nil{
		json.Unmarshal([]byte(val), &s)
		log.Println("进行Redis查询学号")
		return s, nil
	}

	//进行mysql查询
	err = config.DB.QueryRow("SELECT id, name, tel, study FROM students WHERE id = ?", id).
		Scan(&s.ID, &s.Name, &s.Tel, &s.Study)

	if(err != nil){
		return s,err
	}

	data, _ := json.Marshal(s)
    config.RDB.Set(config.Ctx, key, data, 300*time.Second) 
	log.Printf("从 MySQL 中获取了id:%d 学生数据，并写入 Redis 缓存",id)

	return s,err
}

func UpdateStudent(s Student) error{
	result, err := config.DB.Exec("UPDATE students SET name=?, tel=?, study=? WHERE id=?", s.Name, s.Tel, s.Study, s.ID)
    if err != nil{
		return err
	}

	affect, err := result.RowsAffected()

	if err != nil{
		return err
	}

	if affect == 0{
		return errors.New("没有找到对应的学生,或信息未发生改变")
	}


	return nil
}

func DeleteStudent(id int) error{
	result, err := config.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil{
		return err
	}

	affect, err := result.RowsAffected()

	if err != nil{
		return err
	}

	if affect == 0{
		return errors.New("没有查询到对应学生")
	}

	return nil
}