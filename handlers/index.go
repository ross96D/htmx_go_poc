package handlers

import (
	"mpg/htmx_go_poc/webserver/models"

	"github.com/labstack/echo/v4"
)

func HandleIndex(e echo.Context) error {
	tools, err := models.ToolSelectAll(e.Request().Context())
	if err != nil {
		return err
	}
	return e.Render(200, "index.html", map[string]interface{}{
		"body":      "home",
		"Tools":     tools,
		"SearchBar": true,
		"HXRequest": false,
	})
}
