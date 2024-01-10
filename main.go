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

	// Middleware
	e.Static("/uploads", "C:\\Users\\twanm\\reposSchool\\Voicestagram\\Voicemessagefiles")

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("gorm", repositories.Database)
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
	e.GET("/login", handlers.LoginPageRender)
	e.POST("/login/check", handlers.LoginChecker(repositories.Database, store))
	e.GET("/register", handlers.RenderRegistrationPage)
	e.POST("/register", handlers.HandleRegistration(repositories.Database))
	e.GET("/logout", handlers.LogoutHandler(store))
	e.POST("/createpost", handlers.CreatePostHandler)
	e.GET("/createpost", handlers.CreatePostPageRender)
	e.Logger.Fatal(e.Start(":1324"))
}
