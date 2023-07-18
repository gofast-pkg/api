package api

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	portVal := "4242"
	sslEnableVal := "true"
	sslCertVal := "certificates/cert.pem"
	sslKeyVal := "certificates/key.pem"
	swaggerVal := "true"

	t.Run("should return an error when missing port", func(t *testing.T) {
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		a, err := New()
		if assert.Error(t, err) {
			assert.Nil(t, a)
			assert.ErrorIs(t, err, ErrMissingConfig)
		}
	})
	t.Run("should return the api", func(t *testing.T) {
		viper.Set(CONF_HTTP_PORT, portVal)
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		expectedSSL := SSL{
			Enable: true,
			Cert:   sslCertVal,
			Key:    sslKeyVal,
		}
		a, err := New()
		if assert.NoError(t, err) {
			assert.NotNil(t, a)
			assert.EqualValues(t, expectedSSL, a.(*api).ssl)
			assert.EqualValues(t, portVal, a.(*api).port)
		}
	})
}

func TestNewWithConfig(t *testing.T) {
	t.Run("should return the api", func(t *testing.T) {
		c := Config{
			Port:    "8080",
			SSL:     SSL{Enable: true},
			Swagger: true,
		}
		expectedSSL := SSL{Enable: true}

		a, err := NewWithConfig(c)
		if assert.NoError(t, err) {
			assert.NotNil(t, a)
			assert.EqualValues(t, expectedSSL, a.(*api).ssl)
			assert.EqualValues(t, c.Port, a.(*api).port)
		}
	})
}

func TestAPI_SubRouter(t *testing.T) {
	t.Run("should return the subrouter", func(t *testing.T) {
		e := echo.New()
		a := api{driver: e}

		group, err := a.SubRouter("/test", false)

		if assert.NoError(t, err) {
			assert.NotNil(t, group)
		}
	})
}
