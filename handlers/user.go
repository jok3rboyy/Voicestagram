package handlers

import (
	"net/http"

	"github.com/jok3rboyy/VoiceStagram1/repositories"
	"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RenderRegistrationPage renderd de registratie pagina
func RenderRegistrationPage(c echo.Context) error {
	return c.Render(http.StatusOK, "register.gohtml", nil)
}

// HandleRegistration is het handlen van de registratie dngen
func HandleRegistration(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.FormValue("username")
		password := c.FormValue("password")

		// Create een user
		user := types.User{
			Username: username,
		}

		// Sla de user op in de database
		if err := repositories.CreateUser(db, &user, password); err != nil {
			// Handle error
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
		}

		// ga naar login pagina
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}
