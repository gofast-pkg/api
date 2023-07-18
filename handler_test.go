package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Healthcheck(t *testing.T) {
	t.Run("should return the healthcheck", func(t *testing.T) {
		var err error
		var rec *httptest.ResponseRecorder
		var req *http.Request
		var ctx echo.Context

		h := Handler{}

		rec = httptest.NewRecorder()
		req = httptest.NewRequest("GET", healthcheckPath, nil)
		ctx = echo.New().NewContext(req, rec)

		expectedBody := "{\"health\":true}\n"
		err = h.Healthcheck(ctx)
		if assert.NoError(t, err) {
			assert.EqualValues(t, 200, rec.Code)
			assert.EqualValues(t, expectedBody, rec.Body.String())
		}
	})
}
