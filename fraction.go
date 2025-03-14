package main

import "fmt"

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numinator   int // Числитель
	denominator int // Знаменатель
}

func Print(f Fraction){
	fmt.Print(f.numinator)
	fmt.Print("/")
	fmt.Print(f.denominator)
}

func Println(f Fraction){
	fmt.Print(f.numinator)
	fmt.Print("/")
	fmt.Println(f.denominator)
}

func Fraction(a int, b int) Fraction {
	return Fraction{
		numinator   a,
		denominator b,
	}
}

func Fraction(a int, b Fraction) Fraction {
	return Fraction{
		numinator   a * b.denominator,
		denominator b.numinator,
	}
}

func Fraction(a Fraction, b int) Fraction {
	return Fraction{
		numinator   a.numinator,
		denominator b * a.denominator,
	}
}