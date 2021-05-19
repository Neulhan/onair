package database

import (
	"fmt"
	"log"
	"onair/src/model"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetNewConnection(dbName string, con *gorm.Config) (db *gorm.DB) {
	var err error

	for {
		db, err = gorm.Open(postgres.Open(dbName), con)
		if err != nil {
			fmt.Println(dbName)
			fmt.Printf("FAILED -> RECONNECT TO DATABASE[%s]", os.Getenv("MODE"))
			time.Sleep(time.Second * 3)
			continue
		}
		err = db.AutoMigrate(&model.Book{})
		if err != nil {
			log.Fatal("DATABASE MIGRATE FAILED")
		}
		break
	}

	fmt.Println("DATABASE CONNECTED!")
	return
}
