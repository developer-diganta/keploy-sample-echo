package main

import (
	"strings"

	"github.com/keploy/go-sdk/integrations/kecho/v4"
	"github.com/keploy/go-sdk/keploy"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	port:="8080"
	k := keploy.New(keploy.Config{
	App: keploy.AppConfig{
		Name: "my-app",
		Port: port,
	},
	Server: keploy.ServerConfig{
		URL: "http://localhost:6789/api",
	},
	})
e.Use(kecho.EchoMiddlewareV4(k))
	e.GET("/capitalize/:word", func(c echo.Context) error {
		word := c.Param("word")
		return c.String(200, strings.ToUpper(word))
	})

	e.Start(":"+port)
}