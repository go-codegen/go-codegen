package repository_module

import (
	"github.com/go-codegen/go-codegen/internal/constants"
	filesys_core "github.com/go-codegen/go-codegen/internal/filesys/core"
	"github.com/go-codegen/go-codegen/internal/parse"
	"github.com/go-codegen/go-codegen/internal/repository"
	"github.com/go-codegen/go-codegen/internal/utils"
	"strings"
)

type Gorm struct {
	suffix       string
	structSymbol string
	imports      []string
}

func NewGorm() *Gorm {
	return &Gorm{
		suffix:       string(constants.Suffix) + "Impl",
		structSymbol: string(constants.StructSymbol),
	}
}

func (g *Gorm) MethodsData(info parse.ParsedStruct) repository.Methods {
	var methods repository.Methods

	g.addImportFromField("gorm.io/gorm")
	g.addImportFromField(info.PathToPackage)
	methods.Struct = g.createRepositoryStruct(info.StructName + g.suffix)

	methods.Funcs = append(methods.Funcs, g.createFuncNewRepositoryStruct(info.StructName))
	methods.Funcs = append(methods.Funcs, g.create(info))
	methods.Funcs = append(methods.Funcs, g.find(info))
	findFuncs := g.findByAllFields(info)
	for _, f := range findFuncs {
		methods.Funcs = append(methods.Funcs, f)
	}
	methods.Funcs = append(methods.Funcs, g.findByAllFieldsJoin(info)...)

	methods.Funcs = append(methods.Funcs, g.update(info))
	methods.Funcs = append(methods.Funcs, g.delete(info))

	methods.Imports = g.imports

	for _, f := range methods.Funcs {
		if f.Name == "New"+info.StructName+g.suffix {
			continue
		}

		newSuffix := strings.ReplaceAll(g.suffix, "Impl", "")
		methods.Interface.Name = info.StructName + newSuffix

		args := strings.Join(f.Ars, ", ")
		returnValues := strings.Join(f.ReturnValues, ", ")

		stringFunc := f.Name + "(" + args + ") (" + returnValues + ")"

		methods.Interface.Fields = append(methods.Interface.Fields, stringFunc)
	}

	return methods

}

//func NewRepositoryTestRepository(db *gorm.DB) *RepositoryTestRepository {
//	return &RepositoryTestRepository{
//		db: db,
//	}
//}

func (g *Gorm) createFuncNewRepositoryStruct(name string) filesys_core.FuncBody {
	var funcBody filesys_core.FuncBody

	funcBody.Name = "New" + name + g.suffix
	funcBody.Ars = append(funcBody.Ars, "db *gorm.DB")
	funcBody.ReturnValues = append(funcBody.ReturnValues, "*"+name+g.suffix)
	funcBody.Body = "return &" + name + g.suffix + "{\n\t\tdb: db,\t\n\t}"

	return funcBody
}

func (g *Gorm) addImportFromField(imp string) {
	//если импорт уже есть, то не добавляем его
	for _, i := range g.imports {
		if i == imp {
			return
		}
	}

	g.imports = append(g.imports, imp)
}

func (g *Gorm) create(info parse.ParsedStruct) filesys_core.FuncBody {
	var function filesys_core.FuncBody

	entityName := info.StructModule + "." + info.StructName
	variableName := g.getVariableName(info.StructName)

	function.Name = "Create"
	function.StructSymbol = g.structSymbol
	function.StructName = info.StructName + string(g.suffix)

	function.Ars = append(function.Ars, variableName+" *"+entityName)
	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	function.Body = "if err := r.db.Create(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn " + variableName + ", nil"
	return function
}

// getVariableName
func (g *Gorm) getVariableName(name string) string {
	var variableName = strings.ToLower(name[0:1])

	if variableName == g.structSymbol {
		variableName = variableName + "1"
	}

	return variableName
}

func (g *Gorm) createRepositoryStruct(name string) filesys_core.StructBody {
	var entity filesys_core.StructBody

	entity.Name = name
	entity.Fields = append(entity.Fields, "db *gorm.DB")

	return entity
}

// delete func
func (g *Gorm) delete(info parse.ParsedStruct) filesys_core.FuncBody {
	var function filesys_core.FuncBody
	function.Name = "Delete"

	entityName := info.StructModule + "." + info.StructName

	function.StructSymbol = g.structSymbol
	function.StructName = info.StructName + g.suffix
	function.Ars = append(function.Ars, "id string")
	function.ReturnValues = append(function.ReturnValues, "error")

	function.Body = "if err := r.db.Delete(&" + entityName + "{},\"id = ?\", id).Error; err != nil {" + "\n" + "\t\t" + "return err" + "\n" + "\t}" + "\n\n" + "\treturn nil"
	return function
}

// update func
func (g *Gorm) update(info parse.ParsedStruct) filesys_core.FuncBody {
	var function filesys_core.FuncBody

	function.Name = "Update"
	entityName := info.StructModule + "." + info.StructName
	variableName := g.getVariableName(info.StructName)

	function.StructSymbol = g.structSymbol
	function.StructName = info.StructName + string(g.suffix)
	function.Ars = append(function.Ars, variableName+" *"+entityName)
	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	function.Body = "if err := r.db.Save(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn " + variableName + ", nil"
	return function
}

// findById func where id = ?
func (g *Gorm) find(info parse.ParsedStruct) filesys_core.FuncBody {
	var function filesys_core.FuncBody

	function.Name = "FindByID"

	entityName := info.StructModule + "." + info.StructName
	variableName := g.getVariableName(info.StructName)

	function.StructSymbol = "r"
	function.StructName = info.StructName + string(g.suffix)
	function.Ars = append(function.Ars, "id string")
	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	function.Body = "var " + variableName + " " + entityName + "\n\n" + "\tif err := r.db.Where(\"id = ?\", id).First(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn &" + variableName + ", nil"
	return function
}

// // findByAllFields func where field = ?
func (g *Gorm) findByAllFields(info parse.ParsedStruct) []filesys_core.FuncBody {
	var functions []filesys_core.FuncBody

	combinations := g.generateCombinations(info.Fields)

	for _, pf := range combinations {

		unique := true
		skip := false

		for _, f := range pf {
			if skip {
				continue
			}

			checkField := strings.ToLower(f.Name)
			if checkField == "id" || checkField == "gorm.model" || checkField == "createdat" || checkField == "updatedat" || checkField == "deletedat" || checkField == "expiresat" {
				skip = true
			}

			if f.Type == "struct" {
				skip = true
			}

			if index := g.findTag(f.Tags, "index"); index {
				skip = true
			}

			uTag := g.findTag(f.Tags, "unique")
			if !uTag {
				unique = false
			}
		}

		if !skip {
			switch unique {
			case true:
				functions = append(functions, g.findOne(pf, info.StructModule, info.StructName))
			case false:
				functions = append(functions, g.findMany(pf, info.StructModule, info.StructName))
			}
		}
	}

	return functions
}

func (g *Gorm) findTag(tags map[string][]string, tag string) bool {
	res := false
	if tags["gorm"] != nil {
		if len(tags["gorm"]) > 0 {
			res = utils.FindTag(tag, tags["gorm"])
		}
	}
	if tags[string(constants.MainTag)] != nil {
		if len(tags[string(constants.MainTag)]) > 0 {
			res = utils.FindTag(tag, tags[string(constants.MainTag)])
		}
	}

	return res
}

// findBY func where field = ?
func (g *Gorm) findOne(f []parse.ParsedField, packageName, name string) filesys_core.FuncBody {
	var function filesys_core.FuncBody
	entityName := packageName + "." + name
	variableName := g.getVariableName(name)

	function.Name = "FindBy"
	function.StructSymbol = "r"
	function.StructName = name + g.suffix

	for i, field := range f {
		if field.Type == "time.Time" {
			g.addImportFromField("time")
		}
		if field.NestedStruct != nil {
			if field.NestedStruct.PathToPackage != "" {
				g.addImportFromField(field.NestedStruct.PathToPackage)
			}
		}
		if i == 0 {
			function.Name += field.Name
		} else {
			function.Name += "And" + field.Name
		}
		function.Ars = append(function.Ars, field.Name+" "+field.Type)
	}

	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	function.Body = "var " + variableName + " " + entityName + "\n\n" + "\tif err := r.db.Where(\""
	for i, field := range f {
		snakeCase := utils.ParseCamelCaseToSnakeCase(field.Name)
		if i == 0 {
			function.Body += snakeCase + " = ?"
			continue
		}
		function.Body += " AND " + snakeCase + " = ?"
	}
	function.Body += "\", "
	for i, field := range f {
		if i == 0 {
			function.Body += field.Name
			continue
		}

		function.Body += " ," + field.Name
	}

	function.Body += ").First(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn &" + variableName + ", nil"

	return function
}

// findMany func where field = ?
func (g *Gorm) findMany(f []parse.ParsedField, packageName, name string) filesys_core.FuncBody {
	var function filesys_core.FuncBody
	entityName := packageName + "." + name
	variableName := g.getVariableName(name)

	function.Name = "FindBy"
	function.StructSymbol = "r"
	function.StructName = name + g.suffix

	for i, field := range f {
		if field.Type == "time.Time" {
			g.addImportFromField("time")
		}
		if field.NestedStruct != nil {
			if field.NestedStruct.PathToPackage != "" {
				g.addImportFromField(field.NestedStruct.PathToPackage)
			}
		}
		if i == 0 {
			function.Name += field.Name
		} else {
			function.Name += "And" + field.Name
		}
		function.Ars = append(function.Ars, field.Name+" "+field.Type)
	}

	function.ReturnValues = append(function.ReturnValues, "[]*"+entityName, "error")

	function.Body = "var " + variableName + " []*" + entityName + "\n\n" + "\tif err := r.db.Where(\""
	for i, field := range f {
		snakeCase := utils.ParseCamelCaseToSnakeCase(field.Name)
		if i == 0 {
			function.Body += snakeCase + " = ?"
			continue
		}
		function.Body += " AND " + snakeCase + " = ?"
	}
	function.Body += "\", "
	for i, field := range f {
		if i == 0 {
			function.Body += field.Name
			continue
		}

		function.Body += " ," + field.Name
	}

	function.Body += ").Find(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn " + variableName + ", nil"

	return function
}

// join funcs
func (g *Gorm) findByAllFieldsJoin(info parse.ParsedStruct) []filesys_core.FuncBody {
	var functions []filesys_core.FuncBody

	for _, f := range info.Fields {

		checkField := strings.ToLower(f.Name)
		if checkField == "id" || checkField == "gorm.model" || checkField == "createdat" || checkField == "updatedat" || checkField == "deletedat" {
			continue
		}
		if f.Type != string(constants.StructType) {
			continue
		}
		findNestedField := g.findPrefixNameField(info.Fields, f.Name)
		functions = append(functions, g.createJoinFunc(f, findNestedField, info.StructModule, info.StructName)...)
	}

	return functions
}

func (g *Gorm) createJoinFunc(field, nested parse.ParsedField, packageName, name string) []filesys_core.FuncBody {
	var functions []filesys_core.FuncBody
	if field.Type == string(constants.StructType) {
		//Find field where nameHasPrefix struct name

		funcs := g.joinFunc(nested, field.NestedStruct, packageName, name)

		//
		for _, f := range funcs {
			functions = append(functions, f)
		}
	}
	return functions
}

func (g *Gorm) findPrefixNameField(fields []parse.ParsedField, prefix string) parse.ParsedField {
	for _, f := range fields {
		if strings.HasPrefix(f.Name, prefix) {
			return f
		}
	}

	return parse.ParsedField{}
}

func (g *Gorm) joinFunc(field parse.ParsedField, nestedStruct *parse.ParsedStruct, packageName, name string) []filesys_core.FuncBody {
	var functions []filesys_core.FuncBody

	mainStructSnakeCaseName := utils.ParseCamelCaseToSnakeCase(name)
	nestedStructNameCamelCase := utils.ParseCamelCaseToSnakeCase(nestedStruct.StructName)

	//if last symbol is s, add es to end
	if mainStructSnakeCaseName[len(mainStructSnakeCaseName)-1:] == "s" {
		mainStructSnakeCaseName += "es"
	} else {
		mainStructSnakeCaseName += "s"
	}

	if nestedStructNameCamelCase[len(nestedStructNameCamelCase)-1:] == "s" {
		nestedStructNameCamelCase += "es"
	} else {
		nestedStructNameCamelCase += "s"
	}

	nestedField := g.generateCombinations(nestedStruct.Fields)

	for _, f := range nestedField {
		var function filesys_core.FuncBody

		entityName := packageName + "." + name
		variableName := g.getVariableName(name)
		//if last symbol is s, add es to end

		function.Name = "FindBy" + nestedStruct.StructName
		for i, nf := range f {
			if i != len(f)-1 {
				function.Name += nf.Name + "And"
			} else {
				function.Name += nf.Name
			}
		}
		function.StructSymbol = "r"
		function.StructName = name + g.suffix
		for _, nf := range f {
			function.Ars = append(function.Ars, nf.Name+" "+nf.Type)
		}

		function.ReturnValues = append(function.ReturnValues, "[]*"+entityName, "error")

		function.Body = "var " + variableName + " []*" + entityName + "\n\n"
		function.Body += "\tif err := r.db.Table(\"" + mainStructSnakeCaseName + "\").\n"
		function.Body += "\t\tSelect(\"" + mainStructSnakeCaseName + ".*\").\n"
		function.Body += "\t\tJoins(\"JOIN " + nestedStructNameCamelCase + " ON " + nestedStructNameCamelCase + ".id = " + mainStructSnakeCaseName + "." + utils.ParseCamelCaseToSnakeCase(field.Name) + "\").\n"
		function.Body += "\t\tWhere(\""
		for i, nf := range f {
			//if is not last element in slice add AND
			if i != len(f)-1 {
				function.Body += nestedStructNameCamelCase + "." + utils.ParseCamelCaseToSnakeCase(nf.Name) + " = ? AND "
			} else {
				function.Body += nestedStructNameCamelCase + "." + utils.ParseCamelCaseToSnakeCase(nf.Name) + " = ?"
			}
		}
		function.Body += "\", "
		for i, nf := range f {
			//if is not last element in slice add AND
			if i != len(f)-1 {
				function.Body += nf.Name + ", "
			} else {
				function.Body += nf.Name
			}
		}
		function.Body += ").\n"
		function.Body += "\t\tFind(&" + variableName + ").Error; err != nil {\n"
		function.Body += "\t\treturn nil, err\n"
		function.Body += "\t}\n\n"
		function.Body += "\treturn " + variableName + ", nil"

		functions = append(functions, function)
	}

	return functions
}

func (g *Gorm) generateCombinations(arr []parse.ParsedField) [][]parse.ParsedField {
	var combinations [][]parse.ParsedField

	n := len(arr)
	for i := 1; i <= n; i++ {

		g.generateCombination(arr, &combinations, []parse.ParsedField{}, 0, n, i)
	}

	return combinations
}

func (g *Gorm) generateCombination(arr []parse.ParsedField, combinations *[][]parse.ParsedField, current []parse.ParsedField, start, end, size int) {
	if size == 0 {
		*combinations = append(*combinations, append([]parse.ParsedField{}, current...))
		return
	}

	for i := start; i < end; i++ {
		if arr[i].Type == string(constants.StructType) {
			//	skip nested struct
			continue
		}
		current = append(current, arr[i])

		g.generateCombination(arr, combinations, current, i+1, end, size-1)
		current = current[:len(current)-1]
	}
}
