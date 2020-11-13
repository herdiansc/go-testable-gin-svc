package services

import (
	"fmt"
	"testing"
)

func TestPingService_Pong(t *testing.T) {
	cases := []struct{
		testName string
		expected string
	} {
		{
			testName: "Positive: pong",
			expected: "pong",
		},
	}

	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)
		result, _ := PingService.Pong()

		if result != item.expected {
			t.Errorf("response code should %v got %v", item.expected, result)
		}
	}
}

func TestPingService_Divide(t *testing.T) {
	cases := []struct{
		testName string
		request []byte
		expected int
	} {
		{
			testName: "Positive: 200",
			request: []byte(`{"a":4, "b":2}`),
			expected: 200,
		},
		{
			testName: "Negative: 400 - Failed unmarchal",
			request: []byte(`{"a":"c", "b":2}`),
			expected: 400,
		},
		{
			testName: "Negative: 400 - Failed validation",
			request: []byte(`{"a":2}`),
			expected: 400,
		},
	}

	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)
		result := PingService.Divide(item.request)

		if result.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, result.Code)
		}
	}
}

func TestPingService_Calculate(t *testing.T) {
	cases := []struct{
		testName string
		request []byte
		expected int
	} {
		{
			testName: "Positive: 200 - Addition",
			request: []byte(`{"a":4, "operator":"+", "b":2}`),
			expected: 200,
		},
		{
			testName: "Positive: 200 - Substraction",
			request: []byte(`{"a":4, "operator":"-", "b":2}`),
			expected: 200,
		},		{
			testName: "Positive: 200 - Multiplication",
			request: []byte(`{"a":4, "operator":"*", "b":2}`),
			expected: 200,
		},		{
			testName: "Positive: 200 - Division",
			request: []byte(`{"a":4, "operator":"/", "b":2}`),
			expected: 200,
		},
		{
			testName: "Negative: 400 - Failed unmarchal",
			request: []byte(`{"a":"4", "operator":"+", "b":2}`),
			expected: 400,
		},
		{
			testName: "Negative: 400 - Invalid a",
			request: []byte(`{"operator":"+", "b":2}`),
			expected: 400,
		},
		{
			testName: "Negative: 400 - Invalid B should not be 0 for /",
			request: []byte(`{"a":4, "operator":"/", "b":0}`),
			expected: 400,
		},
	}

	for i, item := range cases {
		fmt.Printf("Testing case %d. %s\n", i+1, item.testName)
		result := PingService.Calculate(item.request)

		if result.Code != item.expected {
			t.Errorf("response code should %v got %v", item.expected, result.Code)
		}
	}
}