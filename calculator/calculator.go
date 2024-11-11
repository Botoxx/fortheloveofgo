package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (
	float64,
	error) {
	if b == 0 {
		return 0, errors.New("divison by zero is weird")
	}

	return a / b, nil
}

func Sqrt(input float64) (
	float64,
	error) {
	if input < 0 {
		return 0, errors.New("root of a negative is complex")
	}

	return math.Sqrt(input), nil
}
