package calculator

import (
	"errors"
)

// Add find sum of given two numbers.
func Add(num1, num2 int) int {
	return num1 + num2
}

// Subtract find difference between first number and second number.
func Subtract(num1, num2 int) int {
	return num1 - num2
}

// Multiply find product of given two numbers.
func Multiply(num1, num2 int) int {
	return num1 * num2
}

// Divide find quiotient of first number by second number.
func Divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Division by Zero Error")
	}
	return num1 / num2, nil
}
