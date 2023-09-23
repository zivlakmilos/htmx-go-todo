package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zivlakmilos/htmx-go-todo/internal/items"
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
	e.Use(middleware.Recover())

	e.Static("/", "public/")

	e.GET("/", items.GetItems)
	e.POST("/todos", items.AddItem)
	e.DELETE("/todos/:id", items.DeleteItem)
	e.PUT("/todos/:id", items.UpdateItem)

	e.Logger.Fatal(e.Start(":8080"))
}
