package drv

import fraction "github.com/arsagyr/SARAGO/fraction"

func (drv1 DRV) SumDRV(drv2 DRV) DRV {
	var ar1 []float32
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append(ar1, drv1.values[i]+drv2.values[j])
			ar2 = append(ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1[j] {
				ar2[i] = ar2[i].SumByFraction(ar2[j])
				ar2[j].SetNumerator(0)
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV(drv)
	return drv
}

func (drv1 DRV) SubtractionDRV(drv2 DRV) DRV {
	var ar1 []float32
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append(ar1, drv1.values[i]-drv2.values[j])
			ar2 = append(ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1[j] {
				ar2[i] = ar2[i].SumByFraction(ar2[j])
				ar2[j].SetNumerator(0)
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV(drv)
	return drv
}

func (drv1 DRV) ProductDRV(drv2 DRV) DRV {
	var ar1 []float32
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append(ar1, drv1.values[i]*drv2.values[j])
			ar2 = append(ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1[j] {
				ar2[i] = ar2[i].SumByFraction(ar2[j])
				ar2[j].SetNumerator(0)
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV(drv)
	return drv
}

func (drv1 DRV) RatioDRV(drv2 DRV) DRV {
	var ar1 []float32
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append(ar1, drv1.values[i]/drv2.values[j])
			ar2 = append(ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1[j] {
				ar2[i] = ar2[i].SumByFraction(ar2[j])
				ar2[j].SetNumerator(0)
			}
		}
	}
	var drv = DRV{
		size:          k,
		probabilities: ar2,
		values:        ar1,
	}
	drv = CleanDRV(drv)
	return drv
}
