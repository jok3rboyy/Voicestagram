package handlers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jok3rboyy/VoiceStagram1/types"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePostPageRender(c echo.Context) error {
	return c.Render(http.StatusOK, "createpost.gohtml", nil)
}
func CreatePostHandler(c echo.Context) error {
	// Retrieve the logged-in user's username from the session
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
	}

	// Check if the session token is present
	if sess.Values["token"] == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not logged in")
	}

	// Retrieve the username from the session
	username, ok := sess.Values["Username"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving username from session")
	}

	// Parse form data
	err = c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error parsing form data")
	}

	// Retrieve the voice message file
	file, _, err := c.Request().FormFile("voiceMessage")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error retrieving voice message file")
	}
	defer file.Close()

	// Create a unique filename based on username and timestamp
	filePath := "C:\\Users\\twanm\\reposSchool\\Voicestagram\\Voicemessagefiles\\" + username + "_" + time.Now().Format("20060102150405") + ".wav"

	// Create or open the file on the server
	outFile, err := os.Create(filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating voice message file")
	}
	defer outFile.Close()

	// Copy the contents of the uploaded file to the newly created file on the server
	_, err = io.Copy(outFile, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving voice message file")
	}

	// Create a new post in the database with the file path
	newPost := types.Post{
		Username:             username,
		VoiceMessage:         filePath, // Store the file path in the database
		VoiceMessageFilePath: filePath, // Also store the file path in a separate field if needed
	}

	db, ok := c.Get("gorm").(*gorm.DB)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving database connection")
	}

	result := db.Create(&newPost)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating post")
	}

	// Redirect to the home page or handle the response as needed
	return c.Redirect(http.StatusSeeOther, "/")
}
