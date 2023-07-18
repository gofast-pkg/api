package api

import (
	"errors"

	errorx "github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	// ErrMissingConfig is the error returned when a config key is missing.
	ErrMissingConfig = errors.New("missing config key")
)

var (
	// ConfigurationPrefix is the prefix for the configuration keys.
	// Default is empty.
	// Example:
	//   - if ConfigurationPrefix is "APP", the configuration key for host will be "APP_PGSQL_HOST".
	// After the call to config(), the ConfigurationPrefix is reset to empty.
	ConfigurationPrefix = ""
)

const (
	CONF_HTTP_PORT    = "HTTP_PORT"
	CONF_HTTP_SSL     = "HTTP_SSL_ENABLE"
	CONF_HTTP_CERT    = "HTTP_SSL_CERT"
	CONF_HTTP_KEY     = "HTTP_SSL_KEY"
	CONF_HTTP_SWAGGER = "HTTP_SWAGGER_ENABLE"
)

// SSL specification app
type SSL struct {
	Cert   string
	Key    string
	Enable bool
}

// Config support to the server
type Config struct {
	Port    string
	SSL     SSL
	Swagger bool
}

func config() (Config, error) {
	var c Config
	var err error
	var prefix string

	if ConfigurationPrefix != "" {
		prefix = ConfigurationPrefix + "_"
		ConfigurationPrefix = ""
	}

	if !viper.IsSet(prefix + CONF_HTTP_PORT) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_HTTP_PORT)
	}
	if !viper.IsSet(prefix + CONF_HTTP_SSL) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_HTTP_SSL)
	}
	if !viper.IsSet(prefix + CONF_HTTP_SWAGGER) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_HTTP_SWAGGER)
	}
	c.Port = viper.GetString(prefix + CONF_HTTP_PORT)
	c.Swagger = viper.GetBool(prefix + CONF_HTTP_SWAGGER)
	c.SSL.Enable = viper.GetBool(prefix + CONF_HTTP_SSL)
	if c.SSL.Enable {
		if !viper.IsSet(prefix + CONF_HTTP_CERT) {
			return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_HTTP_CERT)
		}
		if !viper.IsSet(prefix + CONF_HTTP_KEY) {
			return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_HTTP_KEY)
		}
		c.SSL.Cert = viper.GetString(prefix + CONF_HTTP_CERT)
		c.SSL.Key = viper.GetString(prefix + CONF_HTTP_KEY)
	}

	return c, err
}
