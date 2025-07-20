package fraction

import "fmt"

// Методы для конвертации Fraction в Fraction32 и Fraction64
func (f Fraction) ToFraction32() Fraction32 {
	return Fraction32{
		numerator:   int32(f.numerator),
		denominator: int32(f.denominator),
	}

}
func (f Fraction) ToFraction64() Fraction64 {
	return Fraction64{
		numerator:   int64(f.numerator),
		denominator: int64(f.denominator),
	}

}
func (f Fraction32) ToFraction64() Fraction64 {
	return Fraction64{
		numerator:   int64(f.numerator),
		denominator: int64(f.denominator),
	}

}

// Методы для конвертации дроби в другие форматы Fraction
func (f Fraction) ToFloat64() float64 {
	return float64(f.numerator) / float64(f.denominator)
}
func (f Fraction) ToFloat32() float32 {
	return float32(f.numerator) / float32(f.denominator)
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

// Методы для конвертации дроби в другие форматы Fraction32
func (f Fraction32) ToFloat64() float64 {
	return float64(f.numerator) / float64(f.denominator)
}
func (f Fraction32) ToFloat32() float32 {
	return float32(f.numerator) / float32(f.denominator)
}
func (f Fraction32) ToComplex128() complex128 {
	return complex(float64(f.numerator), float64(f.denominator))
}
func (f Fraction32) ToInt32() int32 {
	return int32(f.numerator / f.denominator)
}

// Метод для конвертации дроби в строковый формат
func (f Fraction32) String() string {
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

// Методы для конвертации дроби в другие форматы Fraction64
func (f Fraction64) ToFloat64() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func (f Fraction64) ToComplex128() complex128 {
	return complex(float64(f.numerator), float64(f.denominator))
}
func (f Fraction64) ToInt64() int64 {
	return int64(f.numerator / f.denominator)
}

// Метод для конвертации дроби в строковый формат
func (f Fraction64) String() string {
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}
