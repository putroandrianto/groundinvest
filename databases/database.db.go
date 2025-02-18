package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConnect() {
	dsn := "root:@tcp(127.0.0.1:3306)/groundinvest?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Connected to Database!")
	}
}
