package services

import (
	"fmt"
	"testing"
)

func TestMainService_Pong(t *testing.T) {
	cases := []struct{
		testName string
		expected int
	} {
		{
			testName: "Positive: pong",
			expected: 200,
		},
	}

	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)
		result := MainService.Pong()

		if result.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, result.Code)
		}
	}
}

func TestMainService_HealthCheck(t *testing.T) {
	cases := []struct{
		testName string
		expected int
	} {
		{
			testName: "Positive",
			expected: 200,
		},
	}

	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)
		result := MainService.HealthCheck()

		if result.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, result.Code)
		}
	}
}

func TestMainService_RegisterAppStartTime(t *testing.T) {
	cases := []struct{
		testName string
		expected int
	} {
		{
			testName: "Positive",
			expected: 200,
		},
	}

	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)
		result := MainService.RegisterAppStartTime()

		if result.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, result.Code)
		}
	}
}
