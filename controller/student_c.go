package controller

import (
    "net/http"
    "strconv"
    "student-service/models"
    "student-service/utils"

    "github.com/gin-gonic/gin"
)

//Handle function 
//获取数据结构
//查询数据库
//处理逻辑
//返回响应


//获取所有学生
func ListStudents(context *gin.Context){
	students, err := models.GetAllStudents()
    if err != nil {
        context.JSON(http.StatusInternalServerError,gin.H{
            "error":err.Error(),
        })
        return
    }
	context.JSON(http.StatusOK, students)
}


//创建学生
func CreateStudent(context *gin.Context){
    var student models.Student
    err := context.ShouldBindJSON(&student)
    if err != nil{
        context.JSON(http.StatusBadRequest,gin.H{
            "error":"invalid JOSN",
        })
        return
    }

    err = models.CreateStudent(&student)//传递指针，优化操作
    if err != nil{
        context.JSON(http.StatusInternalServerError,gin.H{
            "error":err.Error(),
        })
        return
    }

    context.JSON(http.StatusCreated,gin.H{
        "message":"student created",
        "student":student,
    })
}	

//获取单个学生
func GetStudent(context *gin.Context){
    id, err := strconv.Atoi(context.Param("id"))
    if err != nil{
        context.JSON(http.StatusBadRequest, gin.H{
            "error":"Invaild ID", 
        })
        return
    }

    student, err := models.GetStudentByID(id)
    if err != nil{
        context.JSON(http.StatusNotFound,gin.H{
            "error":"Not Fount Student",
        })
        return
    }

    context.JSON(http.StatusOK,student)
}

//更新学生
func UpdateStudent(context *gin.Context){
    id, err := strconv.Atoi(context.Param("id"))
    if err != nil{
        context.JSON(http.StatusBadRequest,gin.H{
            "error":"invalid ID",
        })
        return 
    }

    var student models.Student

    err = context.ShouldBindJSON(&student)
    if err != nil{
        context.JSON(http.StatusBadRequest,gin.H{
            "error":"invalid JSON",
        })
        return 
    }   

    student.ID = id
    err = models.UpdateStudent(id,student)

    if err != nil{
        context.JSON(http.StatusInternalServerError,gin.H{
            "error":err.Error(),
        })
        return 
    }

    //对redis清理
    utils.DeleteStudentCache(id)
    
    context.JSON(http.StatusOK,gin.H{
        "mess":"student Update",
        "student":student,
    })
}

//删除学生
func DeleteStudent(context *gin.Context){
    id, err := strconv.Atoi(context.Param("id"))
    if err != nil{
        context.JSON(http.StatusBadRequest,gin.H{
            "error":"invalid ID",
        })
        return
    }

    err = models.DeleteStudent(id)

    if err != nil{
        context.JSON(http.StatusInternalServerError,gin.H{
            "error":err.Error(),
        })
        return
    }

    //对redis清理
    utils.DeleteStudentCache(id)

    context.JSON(http.StatusOK,gin.H{
        "mess":"Succcessful DELETE",
    })
}
