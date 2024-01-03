package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/jok3rboyy/VoiceStagram1/handlers"
	"github.com/jok3rboyy/VoiceStagram1/repositories"

	//"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.Static("/Foto's", "Foto's")

	repositories.DatabaseOpen()

	e.Renderer = echoview.New(goview.Config{
		Root:         "templates",
		Extension:    ".gohtml",
		DisableCache: true,
		Master:       "/master",
	})
	e.GET("/login", handlers.LoginPageRender)                            // Render the login page when /login is accessed
	e.POST("/login/check", handlers.LoginChecker(repositories.Database)) // Login checker (for form submission)
	e.GET("/", handlers.HomeHandler)
	//e.GET("/makePost", handlers.Gotopost)
	//e.POST("/makePost", handlers.Uploadpost)
	e.Logger.Fatal(e.Start(":1324"))

	// func main() {
	// 	// Connect to the MySQL database on localhost
	// 	dsn := "root:Harvey1994!!@tcp(127.0.0.1:3306)/ip_insta_opdracht?parseTime=true"
	// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 	if err != nil {
	// 		panic("Failed to connect to the database")
	// 	}

	// 	// Perform database migrations
	// 	db.AutoMigrate(&types.User{})
	// 	db.AutoMigrate(&types.Post{})

	// 	db.Create(&types.Post{UserID: 1, Message: "This is a sample message"})

	// 	// Initialize the Echo web framework
	// 	e := echo.New()

	// 	// Set up middleware for form handling
	// 	e.Use(middleware.BodyLimit("2M"))
	// 	e.Use(middleware.Logger())
	// 	e.Use(middleware.Recover())

	// 	// Define API endpoints

	// 	// e.POST("/api/voice-messages", createVoiceMessage)
	// 	// e.GET("/api/voice-messages", getVoiceMessages)
	// 	// e.DELETE("/api/voice-messages/:id", deleteVoiceMessage)

	// 	// Login page
	// 	e.GET("/login", handlers.LoginPageHandler)
	// 	e.POST("/login", handlers.LoginHandler(db))

	// 	// Homepage
	// 	e.GET("/", handlers.HomeHandler)

	// // Start the server
	// e.Start(":8080")
}
