package fraction

import "log"

// Метод сравнения, где дробь f меньше числа a
func (f Fraction) IsSmallerThanINT(input any) bool {
	var b bool = false
	switch a := input.(type) {
	case int:
		b = f.numerator < (a * f.denominator)
	default:
		log.Panic("Error: wrong format")
	}
	return b
}

// Метод сравнения, где дробь f больше числа a
func (f Fraction) IsBiggerThanINT(input any) bool {
	var b bool = false
	switch a := input.(type) {
	case int:
		b = f.numerator > (a * f.denominator)
	default:
		log.Panic("Error: wrong format")
	}
	return b
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

// Для Float32
// Метод сравнения, где дробь f меньше числа a
func (f Fraction32) IsSmallerThanINT(a int) bool {
	b := int32(a)
	return f.numerator < (b * f.denominator)
}

// Метод сравнения, где дробь f больше числа a
func (f Fraction32) IsBiggerThanINT(a int) bool {
	b := int32(a)
	return f.numerator > (b * f.denominator)
}

// Метод сравнения, где дробь f равна числу a
func (f Fraction32) IsEqualToINT(a int) bool {
	b := int32(a)
	return f.numerator == (b * f.denominator)
}

// Метод сравнения, где дробь f1 меньше дроби f2
func (f1 Fraction32) IsSmallerThanFraction(f2 Fraction32) bool {
	var b bool

	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) < (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator < f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 больше дроби f2
func (f1 Fraction32) IsBiggerThanFraction(f2 Fraction32) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) > (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator > f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 равна дроби f2
func (f1 Fraction32) IsEqualToFraction(f2 Fraction32) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) == (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator == f2.numerator)
	}
	return b
}

// Для Float64
// Метод сравнения, где дробь f меньше числа a
func (f Fraction64) IsSmallerThanINT(a int) bool {
	b := int64(a)
	return f.numerator < (b * f.denominator)
}

// Метод сравнения, где дробь f больше числа a
func (f Fraction64) IsBiggerThanINT(a int) bool {
	b := int64(a)
	return f.numerator > (b * f.denominator)
}

// Метод сравнения, где дробь f равна числу a
func (f Fraction64) IsEqualToINT(a int) bool {
	b := int64(a)
	return f.numerator == (b * f.denominator)
}

// Метод сравнения, где дробь f1 меньше дроби f2
func (f1 Fraction64) IsSmallerThanFraction(f2 Fraction64) bool {
	var b bool

	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) < (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator < f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 больше дроби f2
func (f1 Fraction64) IsBiggerThanFraction(f2 Fraction64) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) > (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator > f2.numerator)
	}
	return b
}

// Метод сравнения, где дробь f1 равна дроби f2
func (f1 Fraction64) IsEqualToFraction(f2 Fraction64) bool {
	var b bool
	if f1.denominator != f2.denominator {
		b = ((f1.numerator * f2.denominator) == (f2.numerator * f1.denominator))
	} else {
		b = (f1.numerator == f2.numerator)
	}
	return b
}
