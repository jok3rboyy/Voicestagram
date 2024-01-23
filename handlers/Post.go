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
	// pak sessie
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving session")
	}

	// kijk of er een token is
	if sess.Values["token"] == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not logged in")
	}

	// pak de username uit de sessie
	username, ok := sess.Values["Username"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving username from session")
	}

	// parse de form data
	err = c.Request().ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error parsing form data")
	}

	// pak de voicefile uit de form data
	file, _, err := c.Request().FormFile("voiceMessage")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error retrieving voice message file")
	}
	defer file.Close()

	//  maak een filepath die anders is per user en per tijd
	filePath := username + "_" + time.Now().Format("20060102150405") + ".wav"

	// creeer een file op de server op basis van de filepath
	outFile, err := os.Create("C:\\Users\\twanm\\reposSchool\\Voicestagram\\Voicemessagefiles\\" + filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating voice message file")
	}
	defer outFile.Close()

	// gooi deze content in de file die we net hebben gemaakt
	_, err = io.Copy(outFile, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving voice message file")
	}

	// maak een post aan met de juiste data
	newPost := types.Post{
		Username:             username,
		VoiceMessage:         filePath,
		VoiceMessageFilePath: filePath,
	}

	db, ok := c.Get("gorm").(*gorm.DB)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving database connection")
	}

	result := db.Create(&newPost)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating post")
	}

	// ga terug naar de homepagina
	return c.Redirect(http.StatusSeeOther, "/")
}
