package main

import (
	"flag"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (app *App) OpenDB() error {
	db_port := flag.String("db-port", app.GetEnv("DB_PORT", "3308"), "Database port")
	db_host := flag.String("db-host", app.GetEnv("DB_HOST", "localhost"), "Database host")
	db_user := flag.String("db-user", app.GetEnv("DB_USER", "root"), "Database user")
	db_pass := flag.String("db-password", app.GetEnv("DB_PASS", "root"), "Database password")
	db_name := flag.String("db-name", app.GetEnv("DB_NAME", "library"), "Database name")
	flag.Parse()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", *db_user, *db_pass, *db_host, *db_port, *db_name)

	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	sqlDB, err := dbConnection.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	app.db = sqlDB
	app.gormDB = dbConnection
	return nil
}
