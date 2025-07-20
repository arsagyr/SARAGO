package fraction

import "fmt"

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

// Метод для печати дроби
func (f Fraction32) Print() {
	fmt.Print(f.numerator)
	fmt.Print("/")
	fmt.Print(f.denominator)
}

// Метод для печати дроби с новой строки
func (f Fraction32) Println() {
	fmt.Print(f.numerator)
	fmt.Print("/")
	fmt.Println(f.denominator)
}

// Метод для печати дроби
func (f Fraction64) Print() {
	fmt.Print(f.numerator)
	fmt.Print("/")
	fmt.Print(f.denominator)
}

// Метод для печати дроби с новой строки
func (f Fraction64) Println() {
	fmt.Print(f.numerator)
	fmt.Print("/")
	fmt.Println(f.denominator)
}
