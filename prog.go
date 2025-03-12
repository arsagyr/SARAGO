package main

import (
	"fmt"

	// "/theorver"
)

func main() {
	var ar1 []float32 = []float32{0.25, 0.25, 0.25, 0.25}
	var ar2 []float32 = []float32{1, 2, 2, 1}
	var drv = BuildDRV(4, ar1, ar2)
	PrintDRV(drv)
	drv =sumDRV(drv, drv)
	PrintDRV(drv)
	fmt.Println(ExpectedValue((drv)))

}

// type Fraction struct {
// 	numinator   int
// 	denominator int
// }

// type DRV struct {
// 	size int
// 	// probabilities []Fraction
// 	probabilities []float32
// 	values        []float32
// }

// func buildDRV(n int, ar1 []float32, ar2 []float32) DRV {
// 	var drv = DRV{
// 		size:          n,
// 		probabilities: ar1,
// 		values:        ar2,
// 	}
// 	return drv
// }

// func ExpectedValue(drv DRV) float32 {
// 	var ev float32 = 0
// 	for i := 0; i < drv.size; i++ {
// 		ev = ev + (drv.probabilities[i] * drv.values[i])
// 	}
// 	return ev
// }
