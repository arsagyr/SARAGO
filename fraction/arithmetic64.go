package fraction

import "log"

// Метод для умножения дроби на целое число
func (f Fraction64) MultiplyByInt(input any) Fraction64 {
	var a int64
	switch i := input.(type) {
	case int:
		a = int64(i)
	case int32:
		a = int64(i)
	case int64:
		a = i
	default:
		log.Panic("Error: wrong format")
	}

	return Fraction64{
		numerator:   int64(a) * f.numerator,
		denominator: f.denominator,
	}
}

// Метод для умножения дроби на другую дробь
func (f1 Fraction64) MultiplyByFraction(f2 Fraction64) Fraction64 {
	return Fraction64{
		numerator:   f1.numerator * f2.numerator,
		denominator: f1.denominator * f2.denominator,
	}
}

// Метод для деления дроби на целое число
func (f Fraction64) DivideByInt(a int) Fraction64 {
	return Fraction64{
		numerator:   f.numerator,
		denominator: int64(a) * f.denominator,
	}
}

// Метод для деления дроби на другую дробь
func (f1 Fraction64) DivideByFraction(f2 Fraction64) Fraction64 {
	return Fraction64{
		numerator:   f1.numerator * f2.denominator,
		denominator: f1.denominator * f2.numerator,
	}
}

// Метод для вычитания дроби на целое число
func (f Fraction64) SubtractByInt(a int) Fraction64 {
	return Fraction64{
		numerator:   f.numerator - int64(a)*f.denominator,
		denominator: f.denominator,
	}
}

// Метод для вычитания дроби на другую дробь
func (f1 Fraction64) SubtractByFraction(f2 Fraction64) Fraction64 {
	var a, b int64
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) - (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator - f2.numerator
		b = f1.denominator
	}
	return Fraction64{
		numerator:   a,
		denominator: b,
	}
}

// Метод для суммирования дроби на целое число
func (f Fraction64) SumByInt(a int) Fraction64 {
	return Fraction64{
		numerator:   f.numerator + int64(a)*f.denominator,
		denominator: f.denominator,
	}
}

// Метод для суммирования дроби на другую дробь
func (f1 Fraction64) SumByFraction(f2 Fraction64) Fraction64 {
	var a, b int64
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) + (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction64{
		numerator:   a,
		denominator: b,
	}
}

func SumFractions64(f1 Fraction64, f2 Fraction64) Fraction64 {
	var a, b int64
	if f1.denominator != f2.denominator {
		a = (f1.numerator * f2.denominator) + (f2.numerator * f1.denominator)
		b = f2.denominator * f1.denominator
	} else {
		a = f1.numerator + f2.numerator
		b = f1.denominator
	}
	return Fraction64{
		numerator:   a,
		denominator: b,
	}
}
