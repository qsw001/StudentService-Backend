package models

import (
	"encoding/json"
	"fmt"
	"log"
	"student-service/config"
	"time"
)

type Student struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Tel   string `gorm:"type:varchar(20)" json:"tel"`
	Study string `gorm:"type:varchar(100)" json:"study"`
}

// 获取所有学生
func GetAllStudents() ([]Student, error) {
	//先进行redis查询
	key := "student:all"

	val, err := config.RDB.Get(config.Ctx, key).Result()
	if err == nil {
		var students []Student

		err1 := json.Unmarshal([]byte(val), &students)

		if err1 == nil {
			log.Println("从redis中获取了学生数据")
			return students, err
		}
	}

	/*
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
	*/

	//通过grom进行数据库操作
	var students []Student
	err = config.DB.Find(&students).Error

	//写入redis
	data, _ := json.Marshal(students)
	config.RDB.Set(config.Ctx, key, data, 300*time.Second)

	log.Println("从 MySQL 中获取了所有学生数据，并写入 Redis 缓存")

	return students, nil
}

//创建学生

/*
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
*/

func CreateStudent(s *Student) error {
	result := config.DB.Create(s) //这里grom框架已经实现了使用LastInsertId()函数返回ID，所以可以直接使用

	return result.Error
}

func GetStudentByID(id int) (Student, error) {
	var s Student

	//先进行redis查询
	key := fmt.Sprintf("student:%d", id)
	val, err := config.RDB.Get(config.Ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &s)
		log.Println("进行Redis查询学号")
		return s, nil
	}

	/*进行mysql查询
	err = config.DB.QueryRow("SELECT id, name, tel, study FROM students WHERE id = ?", id).
		Scan(&s.ID, &s.Name, &s.Tel, &s.Study)

	if(err != nil){
		return s,err
	}
	*/

	//通过grom进行数据库操作
	result := config.DB.First(&s, id)
	if result.Error != nil {
		return s, result.Error
	}

	data, _ := json.Marshal(s)
	config.RDB.Set(config.Ctx, key, data, 300*time.Second)
	log.Printf("从 MySQL 中获取了id:%d 学生数据，并写入 Redis 缓存", id)

	return s, err
}

func UpdateStudent(id int, s Student) error {
	/*
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
	*/

	var student Student
	result := config.DB.First(&student, id)
	if result.Error != nil {
		return result.Error
	}

	student.Name = s.Name
	student.Tel = s.Tel
	student.Study = s.Study

	config.DB.Save(&student)
	return result.Error

}

func DeleteStudent(id int) error {
	/*
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
	*/

	result := config.DB.Delete(&Student{}, id)
	return result.Error
}
