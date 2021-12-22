package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/models"
	"github.com/herdiansc/go-testable-gin-svc/services"
	"net/http/httptest"
	"testing"
)

type calculatorServiceMock struct {
	divideFn func(request []byte) models.Response
	calculateFn func(request []byte) models.Response
}

func (mock calculatorServiceMock) Divide(request []byte) models.Response {
	return mock.divideFn(request)
}
func (mock calculatorServiceMock) Calculate(request []byte) models.Response {
	return mock.calculateFn(request)
}

func TestCalculatorController_Divide(t *testing.T) {
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

	svcMock := calculatorServiceMock{}
	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)

		svcMock.divideFn = item.divideFn
		services.CalculatorService = svcMock

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		c.Set("payload", []byte(`{"a":2, "b":2}`))
		CalculatorController.Divide(c)

		if r.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, r.Code)
		}
	}
}

func TestCalculatorController_Calculate(t *testing.T) {
	cases := []struct {
		testName string
		calculateFn   func(request []byte) models.Response
		expected int
	}{
		{
			testName: "Positive: 200",
			calculateFn: func(request []byte) models.Response {
				return models.Response{
					Code: 200,
					Content: "2 + by 2 is 4",
					Message: "success",
				}
			},
			expected: 200,
		},
	}

	svcMock := calculatorServiceMock{}
	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)

		svcMock.calculateFn = item.calculateFn
		services.CalculatorService = svcMock

		r := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(r)
		c.Set("payload", []byte(`{"a":2, "operator":"+" "b":2}`))
		CalculatorController.Calculate(c)

		if r.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, r.Code)
		}
	}
}
