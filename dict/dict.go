package fractionDict

import (
	"fmt"
	fraction "github.com/arsagyr/SARAGO/fraction"
)

type Variable struct {
	name     string
	fraction fraction.Fraction
}

type FractionDict []Variable

func (d *FractionDict) AddVariable(v *Variable) {
	*d = append(*d, *v)
}
func (v Variable) Print(){
	fmt.Print(v.name)
	fmt.Print(" ")
	v.fraction.Print()
	fmt.Print(" ")
}
func (v Variable) Println(){
	fmt.Print(v.name)
	fmt.Print(" ")
	v.fraction.Println()
}
func NewVariable(name string, f fraction.Fraction) *Variable {
	return &Variable{
		name:     name,
		fraction: f,
	}
}
