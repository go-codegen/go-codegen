package filesys

type Import struct {
}

func NewImport() *Import {
	return &Import{}
}

func (i *Import) Create(imports []string) string {
	var importString string
	if len(imports) == 0 {
		return ""
	}
	if len(imports) > 1 {
		importString = "import (\n"
		for _, imp := range imports {
			//	//автоматически добавлять "" к импортам если их нет
			importString += "\t\"" + imp + "\"\n"
		}
		importString += ")\n\n"
	} else {
		//автоматически добавлять "" к импортам если их нет
		if imports[0][0] == '"' {
			importString = "import " + imports[0] + "\n\n"
		} else {
			importString = "import \"" + imports[0] + "\"\n\n"
		}
	}
	return importString
}
