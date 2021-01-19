package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/barismar/planner-rest-api/routes"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckShouldReturnOK(t *testing.T) {
	testRouter := routes.SetupRouter()

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}
