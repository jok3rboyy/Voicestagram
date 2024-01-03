package handlers

import (
	"net/http"

	"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// HomeHandler handles requests to the home page
func HomeHandler(c echo.Context) error {
	db, ok := c.Get("gorm").(*gorm.DB)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving database connection")
	}

	// Get the session
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
	}

	// Check if the session token is present
	if sess.Values["token"] == nil {
		// If the session token is not present, show the not logged in message
		return c.Render(http.StatusOK, "home.gohtml", map[string]interface{}{"IsLoggedIn": false})
	}

	// Retrieve the username from the session
	username, ok := sess.Values["Username"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving username from session")
	}

	// Retrieve voicemessages for the logged-in user
	var voicemessages []types.Post
	result := db.Where("username = ?", username).Find(&voicemessages)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving voicemessages")
	}

	// Display voicemessages on the home page
	return c.Render(http.StatusOK, "home.gohtml", map[string]interface{}{"Voicemessages": voicemessages, "IsLoggedIn": true})
}
