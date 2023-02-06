package main

import (
	"fmt"
	"math"
)

func divide(dividend, divisor float64) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	quotient := 0
	divid := math.Abs(dividend)
	divis := math.Abs(divisor)

	for ; divid >= divis; quotient++ {
		divid -= divis
	}

	negative := (dividend > 0) != (divisor > 0)
	if negative {
		return -quotient
	}

	return quotient
}

func main() {
	fmt.Println(divide(-15, 2))

}
