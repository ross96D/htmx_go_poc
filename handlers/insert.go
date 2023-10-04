package handlers

import (
	"database/sql"
	"fmt"
	"mpg/htmx_go_poc/webserver/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func HandleInsertView(e echo.Context) error {
	if e.Request().Header.Get("HX-Request") == "true" {
		return e.Render(200, "insert_view.html", map[string]interface{}{
			"HXRequest": true,
		})
	} else {
		return e.Render(200, "insert_view_full.html", map[string]interface{}{
			"HXRequest": false,
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
		return errors.WithStack(err)
	}
	tools, err := models.ToolSelectAll(e.Request().Context())
	if err != nil {
		return errors.WithStack(err)
	}
	if e.Request().Header.Get("Hx-Request") != "" {
		e.Response().Header().Set("HX-Push", "/")
		return e.Render(200, "home.html", map[string]interface{}{
			"Tools":     tools,
			"HXRequest": true,
		})
	} else {
		return e.Redirect(301, "/")
	}
}

func HandleDelete(e echo.Context) error {
	id := e.QueryParam("id")
	if id == "" {
		return errors.New("no id")
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("on HandleDelete, Atoi: %w", err)
	}
	err = models.ToolDelete(e.Request().Context(), intID)
	if err != nil {
		return fmt.Errorf("on HandleDelete, ToolDelete: %w", err)
	}
	if e.Request().Header.Get("Hx-Request") != "" {
		return e.HTML(200, "")
	} else {
		return e.Redirect(301, "/")
	}
}
