package dict

import (
	"fmt"
	fraction "github.com/arsagyr/SARAGO/fraction"
	drv "github.com/arsagyr/SARAGO/drv"
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
	fmt.Print(", ")
	v.fraction.Print()
	fmt.Print(" ")
}
func (v Variable) Println(){
	fmt.Print(v.name)
	fmt.Print(", ")
	v.fraction.Println()
}
func (v Variable) GetFraction() fraction.Fraction{
	return v.fraction
}
func NewFractionVariable(name string, f fraction.Fraction) *Variable {
	return &Variable{
		name:     name,
		fraction: f,
	}
}

type DRVVariable struct {
	name     string
	drv drv.DRV
}

type DRVDict []DRVVariable

func (d *DRVDict) AddDRVVariable(drv *DRVVariable) {
	*d = append(*d, *drv)
}
func (v DRVVariable) Print(){
	fmt.Print(v.name)
	fmt.Print(", ")
	v.drv.Print()
	fmt.Print(" ")
}
func (v DRVVariable) Println(){
	fmt.Print(v.name)
	fmt.Print(", ")
	v.drv.Println()
}
func NewDRVVariable(name string, drv drv.DRV) *DRVVariable {
	return &DRVVariable{
		name:     name,
		drv: drv,
	}
}