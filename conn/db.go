package conn

import (
	"fmt"
	"go-auth/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func ConnectDB() {
	conf := config.DB()
	logMode := logger.Info

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)
	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	log.Println("DB connection successful")
	if err != nil {
		panic(err)
	}
	db = dB
}

func DB() *gorm.DB {
	return db
}
