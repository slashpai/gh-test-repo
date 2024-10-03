package calculator

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	type testCase struct {
		scenario    string
		num1        int
		num2        int
		expected    int
		expectedErr bool
	}

	testcases := []testCase{
		{
			scenario: "valid addition",
			num1:     5,
			num2:     4,
			expected: 9,
		},
		{
			scenario: "negative addition",
			num1:     -5,
			num2:     -4,
			expected: -9,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("case %s", tc.scenario), func(t *testing.T) {
			res := Add(tc.num1, tc.num2)
			if res != tc.expected {
				t.Fatalf(`Add: expected %d but got %d`, tc.expected, res)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type testCase struct {
		scenario    string
		num1        int
		num2        int
		expected    int
		expectedErr bool
	}

	testcases := []testCase{
		{
			scenario: "valid diff",
			num1:     5,
			num2:     4,
			expected: 1,
		},
		{
			scenario: "negative diff",
			num1:     -5,
			num2:     -4,
			expected: -1,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("case %s", tc.scenario), func(t *testing.T) {
			res := Subtract(tc.num1, tc.num2)
			if res != tc.expected {
				t.Fatalf(`Subtract: expected %d but got %d`, tc.expected, res)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type testCase struct {
		scenario    string
		num1        int
		num2        int
		expected    int
		expectedErr bool
	}

	testcases := []testCase{
		{
			scenario: "valid multiplication",
			num1:     5,
			num2:     4,
			expected: 20,
		},
		{
			scenario: "negative multiplication 1",
			num1:     -5,
			num2:     4,
			expected: -20,
		},
		{
			scenario: "negative multiplication 2",
			num1:     -5,
			num2:     -4,
			expected: 20,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("case %s", tc.scenario), func(t *testing.T) {
			res := Multiply(tc.num1, tc.num2)
			if res != tc.expected {
				t.Fatalf(`Multiply: expected %d but got %d`, tc.expected, res)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type testCase struct {
		scenario    string
		num1        int
		num2        int
		expected    int
		expectedErr bool
	}

	testcases := []testCase{
		{
			scenario: "valid division",
			num1:     50,
			num2:     10,
			expected: 5,
		},
		{
			scenario: "negative division 1",
			num1:     -50,
			num2:     -10,
			expected: 5,
		},
		{
			scenario: "negative division 2",
			num1:     -50,
			num2:     -10,
			expected: 5,
		},
		{
			scenario:    "zero division",
			num1:        5,
			num2:        0,
			expectedErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("case %s", tc.scenario), func(t *testing.T) {
			res, err := Divide(tc.num1, tc.num2)
			if err == nil && tc.expectedErr {
				t.Fatalf("expected error but got nil for num1: %d num2: %d", tc.num1, tc.num2)
			}
			if res != tc.expected {
				t.Fatalf(`Division: expected %d but got %d`, tc.expected, res)
			}
		})
	}
}
