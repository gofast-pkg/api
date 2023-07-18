package api

// Healthcheck model to describe the response http of the healthcheck
// endpoint.
// @Description return the healthcheck of the service
// @Description Health field is true if the service is healthy
type Healthcheck struct {
	Health bool `json:"health"`
}
