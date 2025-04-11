package drv

import (
	"fmt"
	fraction "github.com/arsagyr/SARAGO/fraction"
)
type DRV struct {
	size int
	probabilities []fraction.Fraction
	values        []float32
}
//Метод создания дискретной случайной величины
func NewDRV(n int, ar1 []float32, ar2 []fraction.Fraction) DRV {
	sumf := fraction.NewFraction(0, 1)
	for i := 0 ; i < n; i++ {
		sumf = sumf.SumByFraction(ar2[i])
		if sumf.IsBiggerThanINT(1) == true {
			panic("The Sum of fractions is bigger then 1")
		}
	}
	if !sumf.IsEqualToINT(1) {
		panic("The Sum of fractions isn't 1")
	}
	var drv = DRV{
		size:          n,
		values:        ar1,
		probabilities: ar2,
	}
	return drv
}
//Метод удаления "пустых" значений (значений, где вероятность равна 0)
func CleanDRV(drv DRV) DRV {
	var ar1 []float32 
	var ar2 []fraction.Fraction
	var n int = 0
	for i := 0; i < drv.size; i++ {
		if drv.probabilities[i].GetNumerator()!=0 {		
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

func (drv DRV) Print(){
	n := drv.size -1
	for i := 0; i < n; i++ {
		fmt.Print(drv.values[i])
		fmt.Print(" - ")
		drv.probabilities[i].Print()
		fmt.Print(", ")
	}
	fmt.Print(drv.values[n])
	fmt.Print(" - ")
	drv.probabilities[n].Print()
}
func (drv DRV) Println(){
	for i := 0; i < drv.size; i++ {
		fmt.Print(drv.values[i])
		fmt.Print(" - ")
		drv.probabilities[i].Println()
	}
}



func ExpectedValueFloat32(drv DRV) float32 {
	var ev float32 = 0
	for i := 0; i < drv.size; i++ {
		ev = ev + (drv.probabilities[i].ToFloat32() * drv.values[i])
	}
	return ev
}

func VariationFloat32(drv DRV) float32 {
	var ev float32 = 0
	var v float32 = 0
	for i := 0; i < drv.size; i++ {
		v = v + (drv.probabilities[i].ToFloat32() * drv.values[i] * drv.values[i])
		ev = ev + (drv.probabilities[i].ToFloat32() * drv.values[i])
	}
	return v - ( ev * ev )
}

func (drv1 DRV) SumDRV(drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] + drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
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

func SubtractionDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] - drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
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

func ProductDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] * drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
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

func RatioDRV(drv1 DRV, drv2 DRV) DRV {
	var ar1 []float32 
	var ar2 []fraction.Fraction
	var k int = 0
	for i := 0; i < drv1.size; i++ {
		for j := 0; j < drv2.size; j++ {
			ar1 = append( ar1, drv1.values[i] / drv2.values[j])
			ar2 = append( ar2, drv1.probabilities[i].MultiplyByFraction(drv2.probabilities[j]))
			k = k + 1
		}
	}
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			if ar1[i] == ar1 [j] {
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
