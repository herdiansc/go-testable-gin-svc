package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/models"
	"github.com/herdiansc/go-testable-gin-svc/services"
	"net/http/httptest"
	"testing"
)

type pongServiceMock struct {
	pongFn func() (string, error)
	divideFn func(request []byte) models.Response
}

func (mock pongServiceMock) Pong() (string, error) {
	return mock.pongFn()
}
func (mock pongServiceMock) Divide(request []byte) models.Response {
	return mock.divideFn(request)
}

func TestPingController_Pong(t *testing.T) {
	cases := []struct {
		testName string
		pongFn   func() (string, error)
		expected int
	}{
		{
			testName: "Positive: 200",
			pongFn: func() (string, error) {
				return "pong", nil
			},
			expected: 200,
		},
		{
			testName: "Negative: 500",
			pongFn: func() (string, error) {
				return "", errors.New("error pong")
			},
			expected: 500,
		},
	}

	svcMock := pongServiceMock{}
	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)

		svcMock.pongFn = item.pongFn
		services.PingService = svcMock

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		PingController.Pong(c)

		if r.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, r.Code)
		}
	}
}

func TestPingController_Divide(t *testing.T) {
	cases := []struct {
		testName string
		divideFn   func(request []byte) models.Response
		expected int
	}{
		{
			testName: "Positive: 200",
			divideFn: func(request []byte) models.Response {
				return models.Response{
					Code: 200,
					Content: "4 divided by 4 is 1",
					Message: "success",
				}
			},
			expected: 200,
		},
	}

	svcMock := pongServiceMock{}
	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)

		svcMock.divideFn = item.divideFn
		services.PingService = svcMock

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		c.Set("payload", []byte(`{"a":2, "b":2}`))
		PingController.Divide(c)

		if r.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, r.Code)
		}
	}
}

