package handlers

import (
	"mpg/htmx_go_poc/webserver/models"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func HandleIndex(e echo.Context) error {
	tools, err := models.ToolSelectAll(e.Request().Context())
	if err != nil {
		return errors.WithStack(err)
	}
	return e.Render(200, "home_full.html", map[string]interface{}{
		"Tmp":       "t1.html",
		"Tools":     tools,
		"SearchBar": true,
		"HXRequest": false,
	})
}
