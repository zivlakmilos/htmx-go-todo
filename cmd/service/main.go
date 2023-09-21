package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	tmpls, err := template.New("").ParseGlob("views/*.html")

	if err != nil {
		log.Fatalf("couldn't initialize templates: %v", err)
	}

	e := echo.New()
	e.Renderer = &Template{
		templates: tmpls,
	}

	e.Use(middleware.Logger())
	e.Static("/", "public/")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
