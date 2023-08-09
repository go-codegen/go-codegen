package filesys

type FileBody struct {
	Name       string
	Package    string
	Imports    []string
	Funcs      []FuncBody
	Structs    []StructBody
	Interfaces []InterfaceBody
}

type Body struct {
	function   *Func
	imports    *Import
	structs    *Struct
	interfaces *Interface
}

func NewBody(Func *Func, Import *Import, Struct *Struct, Interface *Interface) *Body {
	return &Body{
		function:   Func,
		imports:    Import,
		structs:    Struct,
		interfaces: Interface,
	}
}

func (b *Body) Create(body FileBody) string {
	var fileBody string

	if body.Package != "" {
		fileBody += "package " + body.Package + "\n\n"
	} else {
		fileBody += "package main\n\n"
	}

	fileBody += b.imports.Create(body.Imports)

	for _, s := range body.Structs {
		fileBody += b.structs.Create(s)
	}

	for _, i := range body.Interfaces {
		fileBody += b.interfaces.Create(i)
	}

	for _, fucns := range body.Funcs {
		fileBody += b.function.Create(fucns)
	}

	return fileBody

}
