package handlers

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Tool struct {
	Name  string
	Date  time.Time
	Color string
	Size  int
}

func HandleInsertView(e echo.Context) error {
	return e.Render(200, "insert_view.html", nil)
}

func HandleInsert(e echo.Context) error {
	date, _ := time.Parse("2006-01-02", e.Request().FormValue("date"))
	size, _ := strconv.Atoi(e.Request().FormValue("size"))

	t := Tool{
		Name:  e.Request().FormValue("name"),
		Date:  date,
		Color: e.Request().FormValue("color"),
		Size:  size,
	}

	println(t.Name)
	return e.Render(200, "home.html", nil)
}
