package helpers

import "strconv"

func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
