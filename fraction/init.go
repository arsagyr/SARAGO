package fraction

// Метод для создания новой дроби
func NewFraction(a int, b int) Fraction {
	if b == 0 {
		panic("Denominator cannot be zero")
	}
	return Fraction{
		numerator:   a,
		denominator: b,
	}
}

func NewFraction32(a int32, b int32) Fraction32 {
	if b == 0 {
		panic("Denominator cannot be zero")
	}
	return Fraction32{
		numerator:   a,
		denominator: b,
	}
}

func NewFraction64(a int64, b int64) Fraction64 {
	if b == 0 {
		panic("Denominator cannot be zero")
	}
	return Fraction64{
		numerator:   a,
		denominator: b,
	}
}
