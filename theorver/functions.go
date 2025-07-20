package theorver

import (
	"fmt"
	"log"
	"math"

	"github.com/arsagyr/SARAGO/fraction"
)

func toInt64(input any) int64 {
	var result int64
	switch a := input.(type) {
	case int:
		result = int64(a)
	case int32:
		result = int64(a)
	case int64:
		result = int64(a)
	default:
		log.Panic("Error: wrong format")
	}
	return result
}

// Число сочетаний
func Comb(ia any, ib any) int64 {
	a := toInt64(ia)
	b := toInt64(ib)
	var i int64
	var c, f int64
	c = 1
	if b < a-b {
		for i = 1; i <= b; i++ {
			c = c * i
		}
		b = a - b
	} else {
		for i = 1; i <= a-b; i++ {
			c = c * i
		}
	}
	f = c
	c = 1
	for i = b + 1; i <= a; i++ {
		c = c * i
	}
	c = c / f
	return c
}

// Число разупорядочиваний
func Discomb(in int) int64 {
	var f int64
	var i int64
	n := int64(in)
	f = 1
	s := fraction.NewFraction64(1, 1)
	for i = 1; i <= n; i++ {
		f = f * i
		s = fraction.SumFractions64(s, fraction.NewFraction64(int64(math.Pow(-1, float64(i))), f))
	}
	s = s.MultiplyByInt(f)
	res := s.ToInt64()
	return res
}

// Ф-я Бернулли
func Bernulli(n int, k int, p float32) float64 {
	var b float64
	c := float64(Comb(n, k))
	b = c * math.Pow(float64(p), float64(k))
	b = b * math.Pow(float64(1-p), float64(n-k))
	return b
}

func PhiCoef(x any) float64 {
	var result float64
	switch a := x.(type) {
	case float32:
		result = (1 + math.Erf(float64(a)/math.Sqrt2)) / 2
	case float64:
		result = (1 + math.Erf(a/math.Sqrt2)) / 2
	default:
		log.Panic("Error: wrong format")
	}
	return result
}

func Integral_De_Moivre_Laplace_theorem(n int, p float32, a float32, b float32) float64 {
	var ml float64
	var np float32 = float32(n) * p
	var npq float32 = float32(math.Sqrt(float64(1-p) * float64(np)))
	ml = PhiCoef(float32(b-np)/npq) - PhiCoef((a-np)/npq)
	return ml
}

func Local_Fucntion_Moivre_laplace(n int, p float32, ik int) float64 {
	k := float32(ik)
	var ml float64
	var np float32 = float32(n) * p
	var sqrtnpq float64 = float64(math.Sqrt(float64(1-p) * float64(np)))
	fmt.Println(sqrtnpq)
	fmt.Println(float64(k-np) / sqrtnpq)
	fmt.Println(PhiCoef((float64(k-np) / sqrtnpq)))

	ml = PhiCoef(float64(k-np)/sqrtnpq) / sqrtnpq
	return ml
}

// double loc_moivre_laplace(int n, int m, double p)
// {
//     double ml, np{n*p},npq;
//     npq=sqrt((1-p)*np);
//     ml=((1/sqrt(2*M_PI))*pow(M_E,(-(pow(((m-np)/npq), 2))/2)))/npq;
//     return ml;
// }
