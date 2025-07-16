package fraction

// Метод сравнения, где дробь f меньше числа a
func (f Fraction) IsSmallerThanINT(a int) bool {
	return f.numerator < (a * f.denominator)
}

// Метод сравнения, где дробь f больше числа a
func (f Fraction) IsBiggerThanINT(a int) bool {
	return f.numerator > (a * f.denominator)
}

// Метод сравнения, где дробь f равна числу a
func (f Fraction) IsEqualToINT(a int) bool {
	return f.numerator == (a * f.denominator)
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
