package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/gorilla/sessions"
	"github.com/jok3rboyy/VoiceStagram1/handlers"
	"github.com/jok3rboyy/VoiceStagram1/repositories"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Database setup
	repositories.DatabaseOpen()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("gorm", repositories.Database)
			return next(c)
		}
	})

	// Middleware to inject the database connection into the Echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("gorm", repositories.Database) // Assuming "gorm" is the key used to store the database connection
			return next(c)
		}
	})

	// Session setup
	store := sessions.NewCookieStore([]byte("secret"))
	e.Use(session.Middleware(store))

	// View setup
	e.Renderer = echoview.New(goview.Config{
		Root:         "templates",
		Extension:    ".gohtml",
		DisableCache: true,
		Master:       "/master",
	})

	e.GET("/", handlers.HomeHandler)
	e.GET("/login", handlers.LoginPageRender)                                   // Render the login page when /login is accessed
	e.POST("/login/check", handlers.LoginChecker(repositories.Database, store)) // Login checker (for form submission)
	e.GET("/register", handlers.RenderRegistrationPage)
	e.POST("/register", handlers.HandleRegistration(repositories.Database))
	e.Logger.Fatal(e.Start(":1324"))
}

//e.GET("/makePost", handlers.Gotopost)
//e.POST("/makePost", handlers.Uploadpost)

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
