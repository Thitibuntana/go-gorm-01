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
	DBUser = "root"
	DBPass = ""
	DBHost = "127.0.0.1"
	DBPort = 3306
	DBName = "store_db"
)

func main() {
	// 1) Connect to MySQL/MariaDB ของ XAMPP และสร้าง store_db หากยังไม่มี
	db, err := connectDB()
	if err != nil {
		log.Fatal("connect database: ", err)
	}
	fmt.Println("Database connection successful")

	// 2) สร้าง/ปรับปรุงตารางจาก Model ทั้ง 4 ตาราง
	if err := db.AutoMigrate(&Category{}, &Product{}, &Customer{}, &Order{}); err != nil {
		log.Fatal("migrate database: ", err)
	}
	fmt.Println("Database migration successful")

	// 3) ตัวอย่างเรียกฟังก์ชัน Read ที่ไม่แก้ไขข้อมูลในฐานข้อมูล
	products, err := getAllProductsDB(db)
	if err != nil {
		log.Println("get all products:", err)
	} else {
		fmt.Printf("All products: %d item(s)\n", len(products))
		for _, product := range products {
			fmt.Printf("ID: %d | Name: %s | Price: %.2f | Stock: %d | Category: %s\n",
				product.ID, product.Name, product.Price, product.Quantity, product.Category.Name)
		}
	}

	// ตัวอย่างเรียกฟังก์ชันอื่น ๆ (เอา // ออกทีละส่วนเมื่อต้องการทดลอง)

	// CREATE
	// category := &Category{Name: "Gaming", Description: "Gaming equipment"}
	// if err := createCategoryDB(db, category); err != nil {
	// 	log.Println("create category:", err)
	// }
	// product := &Product{Name: "Gaming Mouse", Price: 1290, Quantity: 10, CategoryID: category.ID}
	// if err := createProductDB(db, product); err != nil {
	// 	log.Println("create product:", err)
	// }
	// customer := &Customer{Name: "Test Customer", Email: "test@example.com", Phone: "0899999999"}
	// if err := createCustomerDB(db, customer); err != nil {
	// 	log.Println("create customer:", err)
	// }
	// order := &Order{CustomerID: customer.ID, ProductID: product.ID, Quantity: 2, Status: OrderStatusPending}
	// if err := createOrderDB(db, order); err != nil {
	// 	log.Println("create order:", err)
	// }

	// READ / CONDITIONAL QUERY
	// productByID, err := getProductByIDDB(db, 1)
	// productsByName, err := getProductByNameDB(db, "Mouse")
	// productsByCategory, err := getProductsByCategoryDB(db, 3)
	// productsByPrice, err := getProductsByPriceRangeDB(db, 1000, 10000)
	// lowStockProducts, err := getProductsLowStockDB(db, 5)
	// categoryAndPrice, err := getProductsByCategoryAndPriceDB(db, 3, 500, 5000)
	// sortedProducts, err := getProductsSortedByPriceDB(db, "ASC")
	// topProducts, err := getTopExpensiveProductsDB(db, 3)
	// customerByEmail, err := getCustomerByEmailDB(db, "anan@example.com")
	// customerOrders, err := getOrdersByCustomerDB(db, 1)
	// pendingOrders, err := getOrdersByStatusDB(db, OrderStatusPending)
	// ordersByDate, err := getOrdersByDateRangeDB(db, time.Now().AddDate(0, -1, 0), time.Now())
	// ordersWithDetails, err := getOrdersWithDetailsDB(db)
	// ป้องกัน compiler แจ้ง unused variable เมื่อนำตัวอย่างออกมาใช้ ให้ Print ค่าที่ต้องการ เช่น fmt.Println(productByID, err)

	// UPDATE
	// update := &Product{Name: "Updated Product", Price: 1500, Quantity: 20, CategoryID: 3}
	// err = updateProductByIDDB(db, 1, update)
	// err = updateProductPriceDB(db, 1, 1990)
	// err = updateProductStockDB(db, 1, 5) // +5 เพิ่มสต็อก, -5 ลดสต็อก
	// customerUpdate := &Customer{Name: "Updated Customer", Email: "updated@example.com", Phone: "0812345678"}
	// err = updateCustomerByIDDB(db, 1, customerUpdate)
	// err = updateOrderStatusDB(db, 1, OrderStatusCompleted)

	// DELETE / RESTORE
	// err = softDeleteProductDB(db, 1)
	// deletedProducts, err := getDeletedProductsDB(db)
	// err = restoreProductDB(db, 1)
	// err = hardDeleteProductDB(db, 1)
	// err = deleteCustomerByIDDB(db, 1)
	// err = deleteCategoryByIDDB(db, 1)
}

// connectDB creates store_db when necessary, then returns a connection to it.
// The default values match a local XAMPP MySQL/MariaDB installation.
func connectDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	serverDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPass, DBHost, DBPort)
	serverDB, err := gorm.Open(mysql.Open(serverDSN), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	if err := serverDB.Exec("CREATE DATABASE IF NOT EXISTS `" + DBName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci").Error; err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&clientFoundRows=true", DBUser, DBPass, DBHost, DBPort, DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
}
