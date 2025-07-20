package fraction

import "log"

// Метод для умножения дроби на целое число
func (f Fraction32) MultiplyByInt(input any) Fraction32 {
	var a int32
	switch i := input.(type) {
	case int:
		a = int32(i)
	case int32:
		a = i
	default:
		log.Panic("Error: wrong format")
	}

	return Fraction32{
		numerator:   a * f.numerator,
		denominator: f.denominator,
	}
}

// Метод для умножения дроби на другую дробь
func (f1 Fraction32) MultiplyByFraction(f2 Fraction32) Fraction32 {
	return Fraction32{
		numerator:   f1.numerator * f2.numerator,
		denominator: f1.denominator * f2.denominator,
	}
}

// Метод для деления дроби на целое число
func (f Fraction32) DivideByInt(a int) Fraction32 {
	return Fraction32{
		numerator:   f.numerator,
		denominator: int32(a) * f.denominator,
	}
}

// Метод для деления дроби на другую дробь
func (f1 Fraction32) DivideByFraction(f2 Fraction32) Fraction32 {
	return Fraction32{
		numerator:   f1.numerator * f2.denominator,
		denominator: f1.denominator * f2.numerator,
	}
}

// Метод для вычитания дроби на целое число
func (f Fraction32) SubtractByInt(a int) Fraction32 {
	return Fraction32{
		numerator:   f.numerator - int32(a)*f.denominator,
		denominator: f.denominator,
	}
}

// Метод для вычитания дроби на другую дробь
func (f1 Fraction32) SubtractByFraction(f2 Fraction32) Fraction32 {
	var a, b int32
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) - (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator - f2.numerator
		b = f1.denominator
	}
	return Fraction32{
		numerator:   a,
		denominator: b,
	}
}

// Метод для суммирования дроби на целое число
func (f Fraction32) SumByInt(a int) Fraction32 {
	return Fraction32{
		numerator:   f.numerator + int32(a)*f.denominator,
		denominator: f.denominator,
	}
}

// Метод для суммирования дроби на другую дробь
func (f1 Fraction32) SumByFraction(f2 Fraction32) Fraction32 {
	var a, b int32
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) + (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction32{
		numerator:   a,
		denominator: b,
	}
}

func SumFractions32(f1 Fraction32, f2 Fraction32) Fraction32 {
	var a, b int32
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) + (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction32{
		numerator:   a,
		denominator: b,
	}
}
