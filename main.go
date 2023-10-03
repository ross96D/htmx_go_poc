package main

import (
	"io"
	"mpg/htmx_go_poc/handlers"
	"mpg/htmx_go_poc/middlewares"
	"mpg/htmx_go_poc/webserver/db"

	"github.com/philippta/go-template/html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	db.Connect()
	tmpl := template.Must(template.ParseGlob("pages/*.html"))
	template.Must(tmpl.ParseGlob("pages/*/*.html"))

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}
	e.Static("/", "assets")

	e.Use(middlewares.VaryCache)
	e.Use(middlewares.Logger())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.GET("/", handlers.HandleIndex)
	e.GET("/insert_view", handlers.HandleInsertView)
	e.POST("/insert", handlers.HandleInsert)
	e.DELETE("/delete", handlers.HandleDelete)
	e.GET("/delete", handlers.HandleDelete)

	e.Start("localhost:8765")
}
