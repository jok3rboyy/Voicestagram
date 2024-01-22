package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginPageRender(c echo.Context) error {
	return c.Render(http.StatusOK, "login.gohtml", nil)
}

// check login functie
func LoginChecker(db *gorm.DB, store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// pak de user uit de database
		var user types.User
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			fmt.Println("Error retrieving user:", result.Error)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// kijk of de password klopt
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			fmt.Println("Password comparison error:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// maak een session token
		sessionToken := uuid.New().String()

		// pak de session
		session, err := store.Get(c.Request(), "session")
		if err != nil {
			fmt.Println("Error retrieving session:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
		}

		// zet de sessie token en username in de sessie
		session.Values["token"] = sessionToken
		session.Values["Username"] = username

		// sla dat op
		if err := session.Save(c.Request(), c.Response().Writer); err != nil {
			fmt.Println("Error saving session:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error saving session")
		}

		// Redirect naar home
		return c.Redirect(http.StatusSeeOther, "/")
	}
}

// LogoutHandler zorgt dat je ook kan uitloggen
func LogoutHandler(store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		// pak sessie
		session, err := store.Get(c.Request(), "session")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
		}

		// zet de token en username op ""
		session.Values["token"] = ""
		session.Values["Username"] = ""
		session.Options.MaxAge = -1

		// sla de sessie op die leeg is
		if err := session.Save(c.Request(), c.Response().Writer); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error saving session")
		}

		// Redirect naar login
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}
