package utils

import (
	"fmt"
	"student-service/config"
)

func DeleteStudentCache(id int) {
	key := fmt.Sprintf("student:%d", id)
	config.RDB.Del(config.Ctx, key)
}

func DeleteAllStudentsCache() {
	key := "student:all"
	config.RDB.Del(config.Ctx, key)
}