package main

import (
	fraction "github.com/arsagyr/SARAGO/fraction"
	drv "github.com/arsagyr/SARAGO/drv"
	dict "github.com/arsagyr/SARAGO/dict"
	// db "github.com/arsagyr/SARAGO/db"

)

func main() {
	f1 := fraction.NewFraction(1, 4)
	f2 := fraction.NewFraction(3, 4)


	f1.Println() // Вывод: 1/2
	f2.Println() // Вывод: 3/4

	v := dict.NewFractionVariable("A", f2)
	v.Println()

	ar1 := []float32 {0, 1}
	ar2 := []fraction.Fraction {v.GetFraction(), f1}
	drv := drv.NewDRV(2, ar1, ar2)

	v2 := dict.NewDRVVariable("B", drv)
	v2.Print()
	// result := f1.MultiplyByFraction(f2)
	// result.Println() // Вывод: 3/8
}

