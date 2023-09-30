package handlers

import (
	"mpg/htmx_go_poc/webserver/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleInsertView(e echo.Context) error {
	return e.Render(200, "insert_view.html", nil)
}

func HandleInsert(e echo.Context) error {
	date, _ := time.Parse("2006-01-02", e.Request().FormValue("date"))
	size, _ := strconv.Atoi(e.Request().FormValue("size"))

	t := models.Tool{
		Name:  e.Request().FormValue("name"),
		Date:  date,
		Color: e.Request().FormValue("color"),
		Size:  size,
	}

	models.ToolInsert(e.Request().Context(), t)
	return e.Render(200, "home.html", nil)
}
