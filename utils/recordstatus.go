package utils

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user_commit_statu struct {
	Id              int       `gorm:"primaryKey;"`
	Date            time.Time `gorm:"type:date;not null"`
	Dev_name        string    `gorm:"type:varchar(30);not null"`
	Email           string    `gorm:"type:varchar(30);not null"`
	Commit_days     int       `gorm:"type:int;not null"`
	Continuous_days int       `gorm:"type:int;not null"`
}

func RecordStatus(dev, email string, commitCount int, resetContinuesDay bool, dbconf MysqlConf) bool {
	db := connectDB(dbconf)
	user := user_commit_statu{}

	result := db.First(&user, "dev_name = ?", dev)
	if result.Error != nil {
		CreateUserByBranchAndEmail(db, dev, email, time.Now(), 1, 1)
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

func connectDB(dbconf MysqlConf) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconf.User, dbconf.Password, dbconf.Server, dbconf.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Println("failed to connect the database," + err.Error())
	}

	return db
}

func CreateUserByBranchAndEmail(db *gorm.DB, dev, email string, date time.Time, commitDays, continuous_days int) bool {
	user := user_commit_statu{}
	user.Dev_name = dev
	user.Email = email
	user.Date = date
	user.Commit_days = commitDays
	user.Continuous_days = continuous_days

	result := db.Create(&user)
	if result.Error != nil {
		// 处理错误
		log.Println("Error inserting data:", result.Error)
	} else {
		// 插入成功
		log.Println("Success to create user with ID:", user.Dev_name)
	}
	return true
}
