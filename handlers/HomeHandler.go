package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check if the user is logged in by verifying the session token
		session, err := session.Get("session", c)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
		}

		// If the session token is not present, show the not logged in message
		if session.Values["token"] == nil {
			return c.Render(http.StatusOK, "home_not_logged_in.gohtml", nil)
		}

		// Retrieve posts from the database (replace this with your actual logic)
		var posts []types.Post
		result := db.Find(&posts)
		if result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving posts")
		}

		// Display posts on the home page
		return c.Render(http.StatusOK, "home_logged_in.gohtml", map[string]interface{}{"Posts": posts})
	}