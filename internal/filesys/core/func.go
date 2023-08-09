package filesys

type FuncBody struct {
	StructSymbol string
	StructName   string
	Name         string
	Body         string
	Ars          []string
	ReturnValues []string
}

type Func struct {
}

func NewFunc() *Func {
	return &Func{}
}

func (f *Func) Create(body FuncBody) string {
	var funcString string

	if body.StructSymbol != "" {
		funcString += "func (" + body.StructSymbol + " *" + body.StructName + ") "
	} else {
		funcString += "func "
	}

	funcString += body.Name + "("

	if len(body.Ars) > 0 {
		for _, a := range body.Ars {
			funcString += a + ", "
		}
		funcString = funcString[:len(funcString)-2]
	}

	funcString += ")"

	if len(body.ReturnValues) > 0 && len(body.ReturnValues) < 2 {
		funcString += " " + body.ReturnValues[0]
	}
	if len(body.ReturnValues) > 1 {
		funcString += " ("
		for _, r := range body.ReturnValues {
			funcString += r + ", "
		}
		funcString = funcString[:len(funcString)-2]
		funcString += ")"
	}

	funcString += " {\n\t" + body.Body + "\n}\n"

	return funcString
}
