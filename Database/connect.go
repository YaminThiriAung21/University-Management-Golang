package Database

import (
	"fmt"

	"github.com/YaminThiriAung21/UniversityGolang/Model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var err error

func Connectdb() {
	dsn := "root:password@tcp(172.17.0.2:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	Database.AutoMigrate(&Model.Student{})
	Database.AutoMigrate(&Model.Class{})
	fmt.Println("Successfully connected")
}
