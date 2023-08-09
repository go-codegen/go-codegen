package filesys

type InterfaceBody struct {
	Name   string
	Fields []string
}

type Interface struct {
}

func NewInterface() *Interface {
	return &Interface{}
}

func (i *Interface) Create(body InterfaceBody) string {
	var interfaceString string

	interfaceString += "type " + body.Name + " interface {\n"

	for _, f := range body.Fields {
		interfaceString += "\t" + f + "\n"
	}

	interfaceString += "}\n"

	return interfaceString
}
