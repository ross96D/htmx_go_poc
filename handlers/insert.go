package handlers

import (
	"database/sql"
	"mpg/htmx_go_poc/webserver/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleInsertView(e echo.Context) error {
	if e.Request().Header.Get("HX-Request") == "true" {
		return e.Render(200, "insert_view.html", nil)
	} else {
		return e.Render(200, "index.html", map[string]interface{}{
			"body": "insert_view",
		})
	}
}

func HandleInsert(e echo.Context) error {
	date, _ := time.Parse("2006-01-02", e.Request().FormValue("date"))
	size, _ := strconv.Atoi(e.Request().FormValue("size"))

	t := models.Tool{
		Name: e.Request().FormValue("name"),
		Date: sql.NullTime{
			Time:  date,
			Valid: e.Request().FormValue("date") != "",
		},
		Color: e.Request().FormValue("color"),
		Size: sql.NullInt32{
			Int32: int32(size),
			Valid: e.Request().FormValue("size") != "",
		},
	}

	err := models.ToolInsert(e.Request().Context(), t)
	if err != nil {
		return err
	}
	tools, err := models.ToolSelectAll(e.Request().Context())
	if err != nil {
		return err
	}
	e.Response().Header().Set("HX-Push", "/")
	return e.Render(200, "home.html", map[string]interface{}{
		"Tools": tools,
	})
}
