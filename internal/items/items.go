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
	id := c.Param("id")

	for idx, item := range items {
		if item.Id == id {
			if c.FormValue("todo") != "" {
			}
			if c.FormValue("done") == "on" {
				items[idx].Done = true
			} else {
				items[idx].Done = false
			}
			break
		}
	}

	return c.Render(http.StatusOK, "items", items)
}

func DeleteItem(c echo.Context) error {
	id := c.Param("id")

	newItems := make([]Item, 0)
	for _, item := range items {
		if item.Id == id {
			continue
		}
		newItems = append(newItems, item)
	}

	items = newItems

	return c.Render(http.StatusOK, "items", items)
}
