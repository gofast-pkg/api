// Package api: provides a simple way to create a new API
// This package is a wrapper of echo framework.
// https://echo.labstack.com/
package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swag "github.com/swaggo/echo-swagger"
)

// API wrap the echo framwork
// Provides a simple way to create a new API
// Featuring wiht healthcheck and swagger endpoints
type API interface {
	// SubRouter return with the prefix path specified on parameter
	// jwtSecure is a flag to enable jwt middleware
	// the jwtSecure is not implemented yet
	SubRouter(prefix string, jwtSecure bool) (*echo.Group, error)
	// Start run the server on the port specified on config
	Start() error
}

type api struct {
	driver *echo.Echo
	port   string
	ssl    SSL
}

// New instance of API.
// The config is loaded automatically.
func New() (API, error) {
	var c Config
	var err error

	if c, err = config(); err != nil {
		return nil, err
	}

	return NewWithConfig(c)
}

// NewWithConfig instance of API.
// The config is loaded from parameter.
func NewWithConfig(c Config) (API, error) {
	a := api{
		driver: echo.New(),
		ssl:    c.SSL,
		port:   c.Port,
	}
	a.driver.Use(middleware.Logger())
	a.driver.Use(middleware.CORS())

	a.driver.GET(healthcheckPath, Handler{}.Healthcheck)
	if c.Swagger {
		a.driver.GET(swaggerPath, swag.WrapHandler)
	}

	return &a, nil
}

func (a api) SubRouter(prefix string, jwtSecure bool) (*echo.Group, error) {
	router := a.driver.Group(prefix)

	return router, nil
}

func (a api) Start() error {
	var err error

	defer a.driver.Close()
	if a.ssl.Enable {
		if err = a.driver.StartTLS(":"+a.port, a.ssl.Cert, a.ssl.Key); err != nil {
			return err
		}
	} else {
		if err = a.driver.Start(":" + a.port); err != nil {
			return err
		}
	}

	return nil
}
