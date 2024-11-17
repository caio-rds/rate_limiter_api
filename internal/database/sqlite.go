package database

import (
	"fmt"
	"go_limiter_rate/internal/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&schemas.User{}, &schemas.Key{}, &schemas.RequestApi{}, &schemas.Pack{})
	if err != nil {
		return nil
	}

	fmt.Println("Connected to SQLite")
	return db
}
