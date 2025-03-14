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
	drv =SumDRV(drv, drv)
	PrintDRV(drv)
	fmt.Println(ExpectedValue((drv)))

}

