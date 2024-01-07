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

// LoginChecker handles user login
func LoginChecker(db *gorm.DB, store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Retrieve user from the database
		var user types.User
		result := db.Where("username = ?", username).First(&user)
		if result.Error != nil {
			fmt.Println("Error retrieving user:", result.Error)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// Check if the password matches
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			fmt.Println("Password comparison error:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// Generate a session token
		sessionToken := uuid.New().String()

		// Get the session
		session, err := store.Get(c.Request(), "session")
		if err != nil {
			fmt.Println("Error retrieving session:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
		}

		// Set the session token in the session
		session.Values["token"] = sessionToken
		session.Values["Username"] = username

		// Save the session
		if err := session.Save(c.Request(), c.Response().Writer); err != nil {
			fmt.Println("Error saving session:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error saving session")
		}

		// Redirect to the home page
		return c.Redirect(http.StatusSeeOther, "/")
	}
}

// LogoutHandler handles user logout
func LogoutHandler(store *sessions.CookieStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the session
		session, err := store.Get(c.Request(), "session")
		if err != nil {
			// Handle session retrieval error
			return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
		}

		// Clear session data
		session.Values["token"] = ""
		session.Values["Username"] = ""
		session.Options.MaxAge = -1

		// Save the cleared session
		if err := session.Save(c.Request(), c.Response().Writer); err != nil {
			// Handle session save error
			return echo.NewHTTPError(http.StatusInternalServerError, "Error saving session")
		}

		// Redirect to the login page
		return c.Redirect(http.StatusSeeOther, "/login")
	}
}
