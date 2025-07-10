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

func CreateStudent(context *gin.Context){
    var student models.Student
    err := context.ShouldBindJSON(&student)
    if err != nil{
        context.JSON(http.StatusBadRequest,gin.H{
            "error":"invalid JOSN",
        })
    }

    id, err := models.CreateStudent(student)
    if err != nil{
        context.JSON(http.StatusInternalServerError,gin.H{
            "error":err.Error(),
        })
        return
    }

    student.ID = int(id)

    context.JSON(http.StatusCreated,gin.H{
        "message":"student created",
        "student":student,
    })
}	

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
    err = models.UpdateStudent(student)

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
