package main

import (
	"fmt"
	"math"
)

func BinaryToDecimal(binary string) (int, error) {
	decimal := 0
	power := len(binary) - 1

	for _, digit := range binary {
		if digit != '0' && digit != '1' {
			return 0, fmt.Errorf("invalid binary digit: %c", digit)
		}

		if digit == '1' {
			decimal += int(math.Pow(2, float64(power)))
		}
		power--
	}

	return decimal, nil
}
