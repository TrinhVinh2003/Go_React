package database

import (
	"fmt"
	"log"
	"project/vnexpress/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	log.Println("Connection successful.")

	db.AutoMigrate(new(model.Article))

	DBConn = db
}
