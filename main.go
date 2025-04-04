package main

import (
	fraction "github.com/arsagyr/SARAGO/fraction"
	fractionDict "github.com/arsagyr/SARAGO/dict"
)

func main() {
	f1 := fraction.NewFraction(1, 2)
	f2 := fraction.NewFraction(3, 4)


	f1.Println() // Вывод: 1/2
	f2.Println() // Вывод: 3/4

	v := fractionDict.NewVariable("A", f2)
	v.Println()
	result := f1.MultiplyByFraction(f2)
	result.Println() // Вывод: 3/8
}
