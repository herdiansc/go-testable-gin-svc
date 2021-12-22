package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/models"
	"github.com/herdiansc/go-testable-gin-svc/services"
	"net/http/httptest"
	"testing"
)

type mainServiceMock struct {
	pongFn func() models.Response
	healthCheckFn func() models.Response
	registerAppStartTimeFn func() models.Response
}

func (mock mainServiceMock) Pong() models.Response {
	return mock.pongFn()
}
func (mock mainServiceMock) HealthCheck() models.Response {
	return mock.healthCheckFn()
}
func (mock mainServiceMock) RegisterAppStartTime() models.Response {
	return mock.registerAppStartTimeFn()
}

func TestMainController_Pong(t *testing.T) {
	cases := []struct {
		testName string
		pongFn   func() models.Response
		expected int
	}{
		{
			testName: "Positive: 200",
			pongFn: func() models.Response {
				return models.Response{
					Code: 200,
					Content: "pong",
					Message: "success",
				}
			},
			expected: 200,
		},
		{
			testName: "Negative: 500",
			pongFn: func() models.Response {
				return models.Response{
					Code: 500,
					Content: "pong",
					Message: "success",
				}
			},
			expected: 500,
		},
	}

	svcMock := mainServiceMock{}
	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)

		svcMock.pongFn = item.pongFn
		services.MainService = svcMock

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		MainController.Pong(c)

		if r.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, r.Code)
		}
	}
}

func TestMainController_HealthCheck(t *testing.T) {
	cases := []struct {
		testName string
		healthCheckFn   func() models.Response
		expected int
	}{
		{
			testName: "Positive: 200",
			healthCheckFn: func() models.Response {
				return models.Response{
					Code:    200,
					Content: "",
					Message: "success",
				}
			},
			expected: 200,
		},
	}

	svcMock := mainServiceMock{}
	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)

		svcMock.healthCheckFn = item.healthCheckFn
		services.MainService = svcMock

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		MainController.HealthCheck(c)

		if r.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, r.Code)
		}
	}
}