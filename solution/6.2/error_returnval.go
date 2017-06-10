package function

import (
	"errors"
	"math"
)

func ErrorF(num float64) (float64, error) {
	if num < 0 {
		return float64(math.NaN()), errors.New("negative number")
	}
	return math.Sqrt(num), nil
}
func ErrorF2(num float64) (re float64, err error) {
	if num < 0 {
		num, err = float64(math.NaN()), errors.New("negative number")
		return
	}
	num = math.Sqrt(num)
	return
}
