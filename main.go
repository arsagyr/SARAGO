package main

import (
	"fmt"

	"github.com/arsagyr/SARAGO/theorver"
	// db "github.com/arsagyr/SARAGO/db"
)

func main() {
	// f1 := fraction.NewFraction(1, 4)
	// f2 := fraction.NewFraction(3, 4)

	// f1.Println() // Вывод: 1/2
	// f2.Println() // Вывод: 3/4

	// v := dict.NewFractionVariable("A", f2)
	// v.Println()

	// ar1 := []float32 {0, 1}
	// ar2 := []fraction.Fraction {v.GetFraction(), f1}
	// drv := drv.NewDRV(2, ar1, ar2)

	// v2 := dict.NewDRVVariable("B", drv)
	// v2.Print()
	// // result := f1.MultiplyByFraction(f2)
	// // result.Println() // Вывод: 3/8

	fmt.Println(theorver.Integral_De_Moivre_Laplace_theorem(100, 0.8, 85, 90))
	fmt.Println(theorver.Local_Fucntion_Moivre_laplace(700, 0.35, 270))
}
