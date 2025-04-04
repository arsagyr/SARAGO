package fraction

import (
	"fmt"
)

// Структура типа "Fraction" - "Дробь"
type Fraction struct {
	numerator   int // Числитель
	denominator int // Знаменатель
}

// Метод для создания новой дроби
func NewFraction(a int, b int) Fraction {
	//! Добавил проверку на нуль
	if b == 0 {
		panic("Denominator cannot be zero")
	}
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

// Метод для умножения дроби на целое число
func (f Fraction) MultiplyByInt(a int) Fraction {
	return Fraction{
		numerator:   a * f.numerator,
		denominator: f.denominator,
	}
}

// Метод для умножения дроби на другую дробь
func (f1 Fraction) MultiplyByFraction(f2 Fraction) Fraction {
	return Fraction{
		numerator:   f1.numerator * f2.numerator,
		denominator: f1.denominator * f2.denominator,
	}
}

// Метод для деления дроби на целое число
func (f Fraction) DivideByInt(a int) Fraction {
	return Fraction{
		numerator:   f.numerator,
		denominator: a * f.denominator,
	}
}

// Метод для деления дроби на другую дробь
func (f1 Fraction) DivideByFraction(f2 Fraction) Fraction {
	return Fraction{
		numerator:   f1.numerator * f2.denominator,
		denominator: f1.denominator * f2.numerator,
	}
}

// Метод для вычитания дроби на целое число
func (f Fraction) SubtractByInt(a int) Fraction {
	return Fraction{
		numerator:   f.numerator - a*f.denominator,
		denominator: f.denominator,
	}
}

// Метод для вычитания дроби на другую дробь
func (f1 Fraction) SubtractByFraction(f2 Fraction) Fraction {
	var a, b int
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) - (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator - f2.numerator
		b = f1.denominator
	}
	return Fraction{
		numerator:   a,
		denominator: b,
	}
}

// Метод для суммирования дроби на целое число
func (f Fraction) SumByInt(a int) Fraction {
	return Fraction{
		numerator:   f.numerator + a*f.denominator,
		denominator: f.denominator,
	}
}

// Метод для суммирования дроби на другую дробь
func (f1 Fraction) SumByFraction(f2 Fraction) Fraction {
	var a, b int
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) + (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction{
		numerator:   a,
		denominator: b,
	}
}

// Метод сравнения, где дробь f меньше числа a
func (f Fraction) IsSmallerThanINT(a int) bool {
	return f.numerator < a*f.denominator
}

// Метод сравнения, где дробь f больше числа a
func (f Fraction) IsBiggerThanINT(a int) bool {
	return f.numerator > a*f.denominator
}

// Метод сравнения, где дробь f равна числу a
func (f Fraction) IsEqualToINT(a int) bool {
	return f.numerator == a*f.denominator
}

// Метод сравнения, где дробь f1 меньше дроби f2
func (f1 Fraction) IsSmallerThanFraction(f2 Fraction) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) < (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator < f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 больше дроби f2
func (f1 Fraction) IsBiggerThanFraction(f2 Fraction) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) > (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator > f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 равна дроби f2
func (f1 Fraction) IsEqualToFraction(f2 Fraction) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) == (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator == f2.numerator)
	}
	return b
}

// Методы для конвертации дроби в другие форматы
func (f Fraction) ToFloat64() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func (f Fraction) ToComplex128() complex128 {
	return complex(float64(f.numerator), float64(f.denominator))
}

func (f Fraction) ToInt() int {
	return int(f.numerator / f.denominator)
}

// Метод для конвертации дроби в строковый формат
func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}
