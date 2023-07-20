package main

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/karimosman181/simpleGoCRM/database"
	"github.com/karimosman181/simpleGoCRM/lead"
)

/**
 *
 *	create routes
 **/
func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLeads)
	app.Delete(DeleteLeads)
}

/**
 *
 * init DataBase
 **/
func InitDB() {
	var err error

	//open connection
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to DB")
	}

	//migrate data to DB
	database.DBConn.AutoMigrate(&lead.Lead{})
}

/**
 *
 *	main function
 **/
func main() {
	app := fiber.New()

	//inint DB
	InitDB()

	//setup routes
	setupRoutes(app)

	//start server
	app.Listen(3000)

	// close the connection with the DB
	defer database.DBConn.Close()
}
