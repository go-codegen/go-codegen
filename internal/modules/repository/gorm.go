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
}

func NewGorm() *Gorm {
	return &Gorm{
		suffix:       string(constants.Suffix),
		structSymbol: string(constants.StructSymbol),
	}
}

func (g *Gorm) MethodsData(info parse.ParsedStruct) repository.Methods {
	var methods repository.Methods

	methods.Imports = g.createRepositoryImports()

	methods.Struct = g.createRepositoryStruct(info.StructName + string(g.suffix))

	methods.Funcs = append(methods.Funcs, g.create(info))
	methods.Funcs = append(methods.Funcs, g.find(info))

	findFuncs := g.findByAllFields(info)

	for _, f := range findFuncs {
		methods.Funcs = append(methods.Funcs, f)
	}

	methods.Funcs = append(methods.Funcs, g.update(info))
	methods.Funcs = append(methods.Funcs, g.delete(info))

	return methods

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

func (g *Gorm) createRepositoryImports() []string {
	var imports []string

	imports = append(imports, "gorm.io/gorm")

	return imports
}

// delete func
func (g *Gorm) delete(info parse.ParsedStruct) filesys_core.FuncBody {
	var function filesys_core.FuncBody
	function.Name = "Delete"

	entityName := info.StructModule + "." + info.StructName

	function.StructSymbol = g.structSymbol
	function.StructName = info.StructName + string(g.suffix)
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

	for _, f := range info.Fields {
		var function filesys_core.FuncBody

		checkField := strings.ToLower(f.Name)

		if checkField == "id" || checkField == "gorm.model" || checkField == "createdat" || checkField == "updatedat" || checkField == "deletedat" {
			continue
		}

		unique := false
		if f.Tags["gorm"] != nil {
			if len(f.Tags["gorm"]) > 0 {
				unique = utils.FindTag("unique", f.Tags["gorm"])
			}
		}
		if f.Tags[string(constants.MainTag)] != nil {
			if len(f.Tags[string(constants.MainTag)]) > 0 {
				unique = utils.FindTag("unique", f.Tags[string(constants.MainTag)])
			}
		}

		if f.Type == "struct" {
			continue
		}

		if unique {
			//	find one
			function = g.findOne(f, info.StructModule, info.StructName)
			functions = append(functions, function)
		} else {
			//find many
			function = g.findMany(f, info.StructModule, info.StructName)
			functions = append(functions, function)
		}

	}

	return functions
}

// findBY func where field = ?
func (g *Gorm) findOne(f parse.ParsedField, packageName, name string) filesys_core.FuncBody {

	var function filesys_core.FuncBody
	entityName := packageName + "." + name
	variableName := g.getVariableName(name)

	function.Name = "FindBy" + f.Name
	function.StructSymbol = "r"
	function.StructName = name + string(g.suffix)
	function.Ars = append(function.Ars, f.Name+" "+f.Type)
	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	//parse f.Name camel case to snake case
	snakeCase := utils.ParseCamelCaseToSnakeCase(f.Name)

	function.Body = "var " + variableName + " " + entityName + "\n\n" + "\tif err := r.db.Where(\"" + snakeCase + " = ?\", " + f.Name + ").First(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn &" + variableName + ", nil"

	return function
}

// findMany func where field = ?
func (g *Gorm) findMany(f parse.ParsedField, packageName, name string) filesys_core.FuncBody {

	var function filesys_core.FuncBody
	entityName := packageName + "." + name
	variableName := g.getVariableName(name)

	function.Name = "FindBy" + f.Name
	function.StructSymbol = "r"
	function.StructName = name + string(g.suffix)
	function.Ars = append(function.Ars, f.Name+" "+f.Type)
	function.ReturnValues = append(function.ReturnValues, "[]*"+entityName, "error")

	//parse f.Name camel case to snake case
	snakeCase := utils.ParseCamelCaseToSnakeCase(f.Name)

	function.Body = "var " + variableName + " []*" + entityName + "\n\n" + "\tif err := r.db.Where(\"" + snakeCase + " = ?\", " + f.Name + ").Find(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn " + variableName + ", nil"

	return function
}
