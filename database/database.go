package database

import (
	"fmt"

	// "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	dsn := "yourdbuser:yourdbpassword@/demo?parseTime=true"
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
