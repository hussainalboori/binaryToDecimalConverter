package main

import (
	"fmt"
	"math"
	// "strconv"
)

func binaryToDecimal(InReadString string) (int, error) {
	decimal := 0
	power := len(InReadString) - 1

	for _, digit := range InReadString {
		if digit != '0' && digit != '1' {
			return 0, fmt.Errorf("invalid bin value digit: %c", digit)
		}

		if digit == '1' {
			decimal += int(math.Pow(2, float64(power)))
		}
		power--
	}

	return decimal, nil
}