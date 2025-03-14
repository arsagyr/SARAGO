package main

import "fmt"

type DRV struct {
	size int
	// probabilities []Fraction
	values        []float32
	probabilities []float32
}

func BuildDRV(n int, ar1 []float32, ar2 []float32) DRV {
	var drv = DRV{
		size:          n,
		probabilities: ar1,
		values:        ar2,
	}
	return drv
}

func CleanDRV(drv DRV) DRV {
	var ar1 []float32 
	var ar2 []float32 
	var n int = 0
	for i := 0; i < drv.size; i++ {
		
		if drv.probabilities[i] != 0 {		
			ar1 = append(ar1, drv.values[i]) 
			ar2 = append(ar2, drv.probabilities[i])
			n++
		}
	}

	return DRV{
		size:          n,
		values:        ar1,
		probabilities: ar2,
	}
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

func SumDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []float32 
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] + drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i] * drv2.probabilities[j])
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
				ar2[i] = ar2[i] + ar2[j]
				ar2[j] = 0
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV( drv )
	
	return drv
}

func SubtractionDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []float32 
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] - drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i] * drv2.probabilities[j])
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
				ar2[i] = ar2[i] + ar2[j]
				ar2[j] = 0
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV( drv )
	
	return drv
}

func ProductDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []float32 
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] * drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i] * drv2.probabilities[j])
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
				ar2[i] = ar2[i] + ar2[j]
				ar2[j] = 0
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV( drv )
	
	return drv
}

func RatioDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []float32 
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] / drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i] * drv2.probabilities[j])
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
				ar2[i] = ar2[i] + ar2[j]
				ar2[j] = 0
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV( drv )
	
	return drv
}