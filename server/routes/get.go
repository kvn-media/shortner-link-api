package routes

import (
	"net/http"

	"github.com/kvn-media/shortner-link-api/database"
	"github.com/labstack/echo/v4"
)

func GetRedirect(c echo.Context) error {
	linkId := c.QueryParam("id")
	if linkId == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	data := database.GetShortenedLink(linkId)
	if data == nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if data.ProvidedLink != "" {
		return c.Redirect(http.StatusPermanentRedirect, data.ProvidedLink)
	}
	return nil
}
