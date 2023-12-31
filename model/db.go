package model

import (
	"fmt"
	"ginblog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitDb() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.DbUser + ":" + config.DbPassWord + "@tcp(" + config.DbHost + ":" + config.DbPort +
		")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting database", err)
		return
	}

	err = db.AutoMigrate(&User{}, &Category{}, &Article{})
	if err != nil {
		fmt.Println("Error migrating database", err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting database", err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
