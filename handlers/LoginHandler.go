package handlers

import (
	"net/http"
	"time"

	"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginPageRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login.gohtml", nil)
}
func LoginChecker(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// pak user from the database
		var user types.User
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {

			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username ")
		}

		if user.Password != password {

			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password")
		}

		sessionToken := "example-session-token"

		cookie := new(http.Cookie)
		cookie.Name = "session"
		cookie.Value = sessionToken
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		return c.Redirect(http.StatusSeeOther, "/")
	}
}

// 	return e.Render(http.StatusOK, "login", echo.Map{"status": "Plopppir"})
// }

// func LoginChecker(db *gorm.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		username := c.FormValue("username")
// 		password := c.FormValue("password")

// 		// Zoek de gebruiker in de database
// 		var user types.User
// 		result := db.Where("username = ?", username).First(&user)
// 		if result.Error != nil {
// 			return echo.NewHTTPError(401, "Onjuiste gebruikersnaam of wachtwoord")
// 		}

// 		// Controleer het wachtwoord
// 		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 		if err != nil {
// 			return echo.NewHTTPError(401, "Onjuiste gebruikersnaam of wachtwoord")
// 		}

// 		// Inloggen geslaagd
// 		return c.String(200, "Inloggen geslaagd")
// 	}
// }
