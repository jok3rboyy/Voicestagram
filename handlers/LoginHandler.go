package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/twanm/VoiceStagram/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Zoek de gebruiker in de database
		var user types.User
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			return echo.NewHTTPError(401, "Onjuiste gebruikersnaam of wachtwoord")
		}

		// Controleer het wachtwoord
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return echo.NewHTTPError(401, "Onjuiste gebruikersnaam of wachtwoord")
		}

		// Inloggen geslaagd
		return c.String(200, "Inloggen geslaagd")
	}
}
