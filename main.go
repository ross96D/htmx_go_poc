package main

import (
	"html/template"
	"io"
	"log"
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
	tmpls, err := template.New("").ParseGlob("pages/*.html")

	if err != nil {
		log.Fatalf("couldn't initialize templates: %v", err)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: tmpls,
	}
	e.Static("/", "assets")
	loggerConf := middleware.LoggerConfig{
		Format: "${time_rfc3339}\t${remote_ip}\t${method}\t${uri}\t${protocol}\t${status}\t${error}\t${latency}\t${bytes_in}\t${bytes_out}\n",
	}
	e.Use(middleware.LoggerWithConfig(loggerConf))
	e.GET("/", handlers.HandleIndex)
	e.Start("localhost:8765")
}
