package dict

type Variable struct {
	name     string
	fraction Fraction
}

type Dict []Variable

func (d *Dict) AddVariable(v *Variable) {
	*d = append(*d, *v)
}
func NewVariable(name string, f Fraction) *Variable {
	return &Variable{
		name:     name,
		fraction: f,
	}
}
