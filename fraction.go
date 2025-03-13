package main

import "fmt"

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numinator   int // Числитель
	denominator int // Знаменатель
}

func Fraction(a int, b int) Fraction {
	return Fraction{
		numinator   a,
		denominator b,
	}
}


