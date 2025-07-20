package theorver

import (
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
	var a float64
	switch s := x.(type) {
	case int:
		a = float64(s)
	case int32:
		a = float64(s)
	case float32:
		a = float64(s)
	case float64:
		a = s
	default:
		log.Panic("Error: wrong format")
	}
	return (0 + math.Erf(a/math.Sqrt2)) / 2
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
	ml = PhiCoef(float64(k-np)/sqrtnpq) / sqrtnpq
	return ml
}
