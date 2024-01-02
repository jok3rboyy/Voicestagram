package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"message": "Welkom op de homepagina!",
	})
}
