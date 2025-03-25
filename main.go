package main

func main() {
	f1 := NewFraction(1, 2)
	f2 := NewFraction(3, 4)

	f1.Println() // Вывод: 1/2
	f2.Println() // Вывод: 3/4

	result := f1.MultiplyByFraction(f2)
	result.Println() // Вывод: 3/8
}
