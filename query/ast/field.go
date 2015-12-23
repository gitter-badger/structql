package ast

type Field struct {
	Target string
	Name   string
	Value  string
}

func (f *Field) Assemble() string {
	field := ""
	if f.Target != "" {
		field += f.Target + "."
	}
	field += f.Name
	return field
}
