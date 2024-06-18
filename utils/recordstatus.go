package utils

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user_commit_statu struct {
	Id              int       `gorm:"primaryKey;"`
	Date            time.Time `gorm:"type:date;not null"`
	Dev_name        string    `gorm:"type:varchar(10);not null"`
	Email           string    `gorm:"type:varchar(30);not null"`
	Commit_days     int       `gorm:"type:int;not null"`
	Continuous_days int       `gorm:"type:int;not null"`
}

func RecordStatus(dev string, commitCount int, resetContinuesDay bool) bool {
	db := connectDB()
	user := user_commit_statu{}

	result := db.First(&user, "dev_name = ?", dev)
	if result.Error != nil {
		log.Println("Failed to find user", result.Error)
		return false
	}
	//更新用户commit记录
	if resetContinuesDay {
		user.Continuous_days = 0

	} else {
		user.Continuous_days += 1
		user.Commit_days += 1
	}
	user.Date = time.Now()

	result = db.Save(&user)
	if result.Error != nil {
		log.Println("Failed to update user", result.Error)
		return false
	}
	return true
}

func connectDB() *gorm.DB {
	dsn := "root:_@tcp(101.133.169.145:3306)/daily_commit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Println("failed to connect the database," + err.Error())
	}

	return db
}
