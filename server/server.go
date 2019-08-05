package server

import (
	"net/http"

	"github.com/labstack/echo"
)

type server struct {
}

//Start server
func Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
