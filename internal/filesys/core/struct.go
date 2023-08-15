package filesys

import "strings"

type StructFuncBody struct {
	Name         string
	Body         string
	Ars          []string
	ReturnValues []string
}

type StructBody struct {
	Name   string
	Fields []string
	Funcs  []StructFuncBody
}

type Struct struct {
	function *Func
}

func NewStruct(Func *Func) *Struct {
	return &Struct{function: Func}
}

func (s *Struct) Create(body StructBody) string {
	var structString string

	structString += "type " + body.Name + " struct {\n"

	for _, f := range body.Fields {
		structString += "\t" + f + "\n"
	}

	structString += "}\n"

	for _, funcs := range body.Funcs {
		funcStruct := &FuncBody{
			StructSymbol: strings.ToLower(body.Name[0:1]),
			StructName:   body.Name,
			Name:         funcs.Name,
			Body:         funcs.Body,
			Ars:          funcs.Ars,
			ReturnValues: funcs.ReturnValues,
		}
		structString += s.function.Create(*funcStruct) + "\n"
	}

	return structString
}
