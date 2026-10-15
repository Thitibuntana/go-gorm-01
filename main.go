package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	sqlUser     = "root"
	sqlPass     = ""
	sqlHost     = "127.0.0.1"
	sqlPort     = 3306
	databaseName = "store_db"
)

func main() {
	connection, err := openDatabase()
	if err != nil {
		log.Fatal("connect database: ", err)
	}
	fmt.Println("Database connection successful")

	if err := connection.AutoMigrate(&Category{}, &Product{}, &Customer{}, &Order{}); err != nil {
		log.Fatal("migrate database: ", err)
	}
	fmt.Println("Database migration successful")

	inventory, err := fetchAllProducts(connection)
	if err != nil {
		log.Println("get all products:", err)
	} else {
		fmt.Printf("All products: %d item(s)\n", len(inventory))
		for _, item := range inventory {
			fmt.Printf("ID: %d | Name: %s | Price: %.2f | Stock: %d | Category: %s\n",
				item.ID, item.Name, item.Price, item.Quantity, item.Category.Name)
		}
	}
}

func openDatabase() (*gorm.DB, error) {
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	serverDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", sqlUser, sqlPass, sqlHost, sqlPort)
	serverDB, err := gorm.Open(mysql.Open(serverDSN), &gorm.Config{Logger: customLogger})
	if err != nil {
		return nil, err
	}
	if err := serverDB.Exec("CREATE DATABASE IF NOT EXISTS `" + databaseName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci").Error; err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&clientFoundRows=true", sqlUser, sqlPass, sqlHost, sqlPort, databaseName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: customLogger})
}
