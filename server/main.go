package main

import (
	"database/sql"
	"fmt"
	"log"
	"project/vnexpress/crawls"
	"project/vnexpress/database"
	"project/vnexpress/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func init() {
	database.ConnectDB()
	crawls.Crawl()
}
func main() {

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	app.Listen(":8000")
}
