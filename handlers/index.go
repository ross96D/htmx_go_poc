package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HandleIndex(e echo.Context) error {
	err := e.Render(200, "index.html", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
