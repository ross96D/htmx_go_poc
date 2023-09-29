package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string
	Sex  string
	Age  int
}

func HandleIndex(e echo.Context) error {
	cats := []Cat{
		{
			Name: "Muchi",
			Sex:  "Macho",
			Age:  4,
		},
		{
			Name: "Hamburguesa",
			Sex:  "Hembra",
			Age:  2,
		},
		{
			Name: "Cuco",
			Sex:  "Macho",
			Age:  1,
		},
	}
	err := e.Render(200, "index.html", cats)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
