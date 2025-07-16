package fraction

import (
	"fmt"
)

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numerator   int // Числитель
	denominator int // Знаменатель
}

type Fraction32 struct {
	numerator   int32 // Числитель
	denominator int32 // Знаменатель
}

type Fraction64 struct {
	numerator   int64 // Числитель
	denominator int64 // Знаменатель
}

func (f Fraction) GetNumerator() int {
	return f.numerator
}
func (f *Fraction) SetNumerator(a int) {
	f.numerator = a
}
func (f Fraction) GetDenominator() int {
	return f.denominator
}

func (f Fraction32) GetNumerator() int32 {
	return f.numerator
}
func (f *Fraction32) SetNumerator(a int32) {
	f.numerator = a
}
func (f Fraction32) GetDenominator() int32 {
	return f.denominator
}

func (f Fraction64) GetNumerator() int64 {
	return f.numerator
}
func (f *Fraction64) SetNumerator(a int64) {
	f.numerator = a
}
func (f Fraction64) GetDenominator() int64 {
	return f.denominator
}

// Метод для печати дроби
func (f Fraction) Print() {
	fmt.Print(f.numerator)
	fmt.Print("/")
	fmt.Print(f.denominator)
}

// Метод для печати дроби с новой строки
func (f Fraction) Println() {
	fmt.Print(f.numerator)
	fmt.Print("/")
	fmt.Println(f.denominator)
}

// Метод упрощения дроби путем нахождения НОД с помощью алгоритма Евклида
func (f Fraction) Simplification() Fraction {
	var m, n, r int
	var i int = 0
	m = f.numerator
	n = f.denominator
	if m > n {
		i = m / n
		m = m % n
	}
	r = m % n
	for r != 0 {
		m = n
		n = r
		r = m % n
	}
	return Fraction{
		denominator: f.denominator / n,
		numerator:   f.numerator/n + f.denominator*i,
	}
}
