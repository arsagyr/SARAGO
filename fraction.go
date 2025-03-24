package main

import "fmt"

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numerator   int // Числитель
	denominator int // Знаменатель
}

// Метод для создания новой дроби
func NewFraction(a int, b int) Fraction {
	return Fraction{
		numerator:   a,
		denominator: b,
	}
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

// Метод для умножения дроби на целое число
func (f Fraction) MultiplyByInt(a int) Fraction {
	return Fraction{
		numerator:   a * f.numerator,
		denominator: f.denominator,
	}
}

// Метод для умножения дроби на другую дробь
func (f Fraction) MultiplyByFraction(b Fraction) Fraction {
	return Fraction{
		numerator:   f.numerator * b.numerator,
		denominator: f.denominator * b.denominator,
	}
}
