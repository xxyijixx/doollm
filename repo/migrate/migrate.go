package main

import (
	"doollm/config"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := config.EnvConfig.GetDSN()
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: dsn,
			}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		panic(fmt.Errorf("db connection failed: %v", err))
	}
	_ = db
	// err = db.AutoMigrate()
	// if err != nil {
	// 	panic(fmt.Errorf("db migrate failed: %v", err))
	// }
}
