package std

import (
	"fmt"
	"math"
	"strings"
)

func ConvertToMilliseconds(value float64, unit string) (float64, error) {
	var err error
	var result float64

	result = value

	switch strings.ToLower(unit) {
	case "µs", "us":
		result = value / 1000
	case "ms":
		result = value
	case "sec":
		result = value * 1000
	case "min":
		result = value * 1000 * 60
	case "hr":
		result = value * 1000 * 60 * 60
	case "day":
		result = value * 1000 * 60 * 60 * 24
	default:
		err = fmt.Errorf("unknown unit: %s", unit)
	}

	return result, err
}

func ConvertToBytes(value float64, unit string) (float64, error) {
	var err error
	var result float64

	result = value

	switch strings.ToUpper(unit) {
	case "BYTES", "B":
		result = value
	case "KB":
		result = value * 1024
	case "MB":
		result = value * 1024 * 1024
	case "GB":
		result = value * 1024 * 1024 * 1024
	case "TB":
		result = value * 1024 * 1024 * 1024 * 1024
	default:
		err = fmt.Errorf("unknown unit: %s", unit)
	}

	return result, err
}

func ConvertToNumeric(value float64, unit string) (float64, error) {
	var err error
	var result float64

	result = value

	switch strings.ToLower(unit) {
	case "":
		result = value
	case "hundred":
		result = value * 100
	case "thousand":
		result = value * 1000
	case "lakh":
		result = value * 100000
	case "million":
		result = value * 1000000
	case "billion":
		result = value * 1000000000
	case "trillion":
		result = value * 1000000000000
	default:
		err = fmt.Errorf("unknown unit: %s", unit)
	}

	return result, err
}

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func Float64ComparisonWithTolerance(a float64, b float64, tolerance float64) bool {
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	}
	return false
}
