package server

import (
	"net/http"
	"os"
	"strconv"

	"../../app/config"
	"github.com/labstack/echo"
)

//Start server
func Start() {
	e := echo.New()
	wd, _ := os.Getwd()
	cfg, configErr := config.LoadEnvConfig(wd + "/app/config/config.json")
	if configErr != nil {
		e.Logger.Fatal("Can't start server, configuration not found")
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start("localhost:" + strconv.Itoa(cfg.Port)))
}
