package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Item struct {
	Id    string
	Title string
	Done  bool
}

var items = make([]Item, 0)

func GetItems(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", items)
}

func AddItem(c echo.Context) error {
	items = append(items, Item{
		Id:    uuid.NewString(),
		Title: c.FormValue("todo"),
		Done:  false,
	})

	return c.Render(http.StatusOK, "items", items)
}

func UpdateItem(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", items)
}

func DeleteItem(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", items)
}
