package api

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const prefixConfiguration = "API"

type API interface {
	SubRouter(prefix string, jwtSecure bool) (*echo.Group, error)
	Start() error
}

// api component
type api struct {
	driver *echo.Echo
	ssl    SSL
	port   string
}

// New app http
func New() (API, error) {
	var c Config
	var err error

	if err = envconfig.Process("", &c); err != nil {
		return nil, err
	}
	return NewWithConfig(c)
}

// NewWithConfig server http
func NewWithConfig(c Config) (API, error) {
	a := api{
		driver: echo.New(),
		ssl:    c.Ssl,
		port:   c.Port,
	}
	a.driver.Use(middleware.Logger())
	a.driver.Use(middleware.CORS())

	a.driver.GET("/ping", Ping)
	return a, nil
}

// SubRouter return with the prefix path specified on parameter
func (a api) SubRouter(prefix string, jwtSecure bool) (*echo.Group, error) {
	router := a.driver.Group(prefix)

	return router, nil
}

// Start run the server
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
