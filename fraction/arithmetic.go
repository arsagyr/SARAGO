package fraction

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

func SumFractions(f1 Fraction, f2 Fraction) Fraction {
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
