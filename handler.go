package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	healthcheckPath = "/healthcheck"
	swaggerPath     = "/swagger/*"
)

type Handler struct{}

// Healthcheck method
// @Summary Healthcheck
// @Description return the healthcheck of the service
// @Tags healthcheck
// @Produce json
// @Success 200 {object} Healthcheck
// @Router /healthcheck [get]
func (h Handler) Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, &Healthcheck{Health: true})
}
