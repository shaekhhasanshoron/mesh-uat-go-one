package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"mesh-uat-go-one/config"
	"mesh-uat-go-one/router"
	"mesh-uat-go-one/server"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	config.InitEnvironmentVariables()

	srv := server.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	srv.Renderer = renderer
	router.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
