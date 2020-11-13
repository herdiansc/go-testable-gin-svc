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