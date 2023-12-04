package main

import (
	"fmt"
	"go-project/database"
	"go-project/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "go-project/models"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.UserProfile{})
}

func main() {
	app := fiber.New()
	routes.Routes(app)
	routes.InetRoutes(app)
	routes.Routes(app)
	initDatabase()
	app.Listen(":3000")
}
