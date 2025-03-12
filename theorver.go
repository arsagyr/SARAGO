package main

import "fmt"

type DRV struct {
	size int
	// probabilities []Fraction
	probabilities []float32
	values        []float32
}

func BuildDRV(n int, ar1 []float32, ar2 []float32) DRV {
	var drv = DRV{
		size:          n,
		probabilities: ar1,
		values:        ar2,
	}
	return drv
}

func PrintDRV(drv DRV) {
	for i := 0; i < drv.size; i++ {
		fmt.Print(drv.values[i])
		fmt.Print(" - ")
		fmt.Println(drv.probabilities[i])
	}
	
}

func ExpectedValue(drv DRV) float32 {
	var ev float32 = 0
	for i := 0; i < drv.size; i++ {
		ev = ev + (drv.probabilities[i] * drv.values[i])
	}
	return ev
}

func Variation(drv DRV) float32 {
	var ev float32 = 0
	var v float32 = 0
	for i := 0; i < drv.size; i++ {
		v = v + (drv.probabilities[i] * drv.values[i] * drv.values[i])
		ev = ev + (drv.probabilities[i] * drv.values[i])
	}
	return v - ( ev * ev )
}

func sumDRV(drv1 DRV, drv2 DRV) DRV {

	ar1 []float32  drv1.size * drv2.size
	ar2 []float32 drv1.size * drv2.size

	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1[k] = drv1.values[i] + drv2.values[j]
			ar2[k] = drv1.probabilities[i] * drv2.probabilities[j]
			k = k + 1
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar1,
		values:        ar2,
	}
	
	return drv
}