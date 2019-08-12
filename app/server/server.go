package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../../app/config"
	"github.com/labstack/echo"
)

//Start server
func Start() {
	e := echo.New()
	cfg, configErr := config.LoadEnvConfig("../config/config.json")
	fmt.Println(configErr)
	if configErr != nil {
		log.Fatal("Can't start server, configuration not found")
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.Port)))
}
