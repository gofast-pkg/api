package api

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	portVal := "4242"
	sslEnableVal := "true"
	sslCertVal := "certificates/cert.pem"
	sslKeyVal := "certificates/key.pem"
	swaggerVal := "true"

	t.Run("should return the config with prefix", func(t *testing.T) {
		prefix := "APP"
		ConfigurationPrefix = prefix

		prefix += "_"
		viper.Set(prefix+CONF_HTTP_PORT, portVal)
		viper.Set(prefix+CONF_HTTP_SSL, sslEnableVal)
		viper.Set(prefix+CONF_HTTP_CERT, sslCertVal)
		viper.Set(prefix+CONF_HTTP_KEY, sslKeyVal)
		viper.Set(prefix+CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		config, err := config()
		if assert.NoError(t, err) {
			assert.EqualValues(t, portVal, config.Port)
			assert.True(t, config.SSL.Enable)
			assert.EqualValues(t, sslCertVal, config.SSL.Cert)
			assert.EqualValues(t, sslKeyVal, config.SSL.Key)
			assert.True(t, config.Swagger)
		}
	})
	t.Run("should return the config without prefix", func(t *testing.T) {
		viper.Set(CONF_HTTP_PORT, portVal)
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		config, err := config()
		if assert.NoError(t, err) {
			assert.EqualValues(t, portVal, config.Port)
			assert.True(t, config.SSL.Enable)
			assert.EqualValues(t, sslCertVal, config.SSL.Cert)
			assert.EqualValues(t, sslKeyVal, config.SSL.Key)
			assert.True(t, config.Swagger)
		}
	})
	t.Run("should return an error when missing port", func(t *testing.T) {
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		config, err := config()
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMissingConfig)
			assert.Empty(t, config)
		}
	})
	t.Run("should return an error when missing ssl", func(t *testing.T) {
		viper.Set(CONF_HTTP_PORT, portVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		config, err := config()
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMissingConfig)
			assert.Empty(t, config)
		}
	})
	t.Run("should return an error when missing cert", func(t *testing.T) {
		viper.Set(CONF_HTTP_PORT, portVal)
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		config, err := config()
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMissingConfig)
			assert.Empty(t, config)
		}
	})
	t.Run("should return an error when missing key", func(t *testing.T) {
		viper.Set(CONF_HTTP_PORT, portVal)
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_SWAGGER, swaggerVal)
		defer viper.Reset()

		config, err := config()
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMissingConfig)
			assert.Empty(t, config)
		}
	})
	t.Run("should return an error when missing swagger", func(t *testing.T) {
		viper.Set(CONF_HTTP_PORT, portVal)
		viper.Set(CONF_HTTP_SSL, sslEnableVal)
		viper.Set(CONF_HTTP_CERT, sslCertVal)
		viper.Set(CONF_HTTP_KEY, sslKeyVal)
		defer viper.Reset()

		config, err := config()
		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrMissingConfig)
			assert.Empty(t, config)
		}
	})
}
