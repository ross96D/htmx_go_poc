package main

import (
	"html/template"
	"io"
	"mpg/htmx_go_poc/handlers"

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
	tmpl := template.Must(template.ParseGlob("pages/*.html"))
	template.Must(tmpl.ParseGlob("pages/*/*.html"))

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}
	e.Static("/", "assets")
	loggerConf := middleware.LoggerConfig{
		Format: "${time_rfc3339} ip:${remote_ip} method:${method} uri:${uri} ${protocol} scode:${status} e:${error} ms:${latency} in:${bytes_in} out:${bytes_out}\n",
	}
	e.Use(middleware.LoggerWithConfig(loggerConf))

	e.GET("/", handlers.HandleIndex)
	e.GET("/insert_view", handlers.HandleInsertView)
	e.POST("/insert", handlers.HandleInsert)

	e.Start("localhost:8765")
}
