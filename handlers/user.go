package handlers

import (
	"net/http"

	"github.com/jok3rboyy/VoiceStagram1/repositories"
	"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RenderRegistrationPage renders the registration page
func RenderRegistrationPage(c echo.Context) error {
	return c.Render(http.StatusOK, "register.gohtml", nil)
}

// HandleRegistration handles the registration process
func HandleRegistration(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get form values
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Your validation logic...

		// Create a new user
		user := types.User{
			Username: username,
		}

		// Save the user to the database (including password hashing)
		if err := repositories.CreateUser(db, &user, password); err != nil {
			// Handle error...
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
		}

		// Redirect to the login page or any other page
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}
