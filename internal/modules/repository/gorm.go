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

func (g *Gorm) MethodsData(info parse.StructInfo) repository.Methods {
	var methods repository.Methods

	methods.Imports = g.createRepositoryImports()

	methods.Struct = g.createRepositoryStruct(info.Name + string(g.suffix))

	methods.Funcs = append(methods.Funcs, g.create(info))
	methods.Funcs = append(methods.Funcs, g.find(info))
	findFuncs := g.findByAllFields(info)

	for _, f := range findFuncs {
		methods.Funcs = append(methods.Funcs, f)
	}
	methods.Funcs = append(methods.Funcs, g.update(info))
	methods.Funcs = append(methods.Funcs, g.delete(info))

	for _, f := range methods.Funcs {
		methods.Interface.Name = info.Name + string(g.suffix) + "Impl"

		args := strings.Join(f.Ars, ", ")
		returnValues := strings.Join(f.ReturnValues, ", ")

		stringFunc := f.Name + "(" + args + ") (" + returnValues + ")"

		methods.Interface.Fields = append(methods.Interface.Fields, stringFunc)
	}

	return methods

}

func (g *Gorm) create(info parse.StructInfo) filesys_core.FuncBody {
	var function filesys_core.FuncBody

	entityName := info.PackageName + "." + info.Name
	variableName := g.getVariableName(info.Name)

	function.Name = "Create"
	function.StructSymbol = g.structSymbol
	function.StructName = info.Name + string(g.suffix)

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
func (g *Gorm) delete(info parse.StructInfo) filesys_core.FuncBody {
	var function filesys_core.FuncBody
	function.Name = "Delete"

	entityName := info.PackageName + "." + info.Name

	function.StructSymbol = g.structSymbol
	function.StructName = info.Name + string(g.suffix)
	function.Ars = append(function.Ars, "id string")
	function.ReturnValues = append(function.ReturnValues, "error")

	function.Body = "if err := r.db.Delete(&" + entityName + "{},\"id = ?\", id).Error; err != nil {" + "\n" + "\t\t" + "return err" + "\n" + "\t}" + "\n\n" + "\treturn nil"
	return function
}

// update func
func (g *Gorm) update(info parse.StructInfo) filesys_core.FuncBody {
	var function filesys_core.FuncBody

	function.Name = "Update"
	entityName := info.PackageName + "." + info.Name
	variableName := g.getVariableName(info.Name)

	function.StructSymbol = g.structSymbol
	function.StructName = info.Name + string(g.suffix)
	function.Ars = append(function.Ars, variableName+" *"+entityName)
	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	function.Body = "if err := r.db.Save(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn " + variableName + ", nil"
	return function
}

// findById func where id = ?
func (g *Gorm) find(info parse.StructInfo) filesys_core.FuncBody {
	var function filesys_core.FuncBody

	function.Name = "FindByID"

	entityName := info.PackageName + "." + info.Name
	variableName := g.getVariableName(info.Name)

	function.StructSymbol = "r"
	function.StructName = info.Name + string(g.suffix)
	function.Ars = append(function.Ars, "id string")
	function.ReturnValues = append(function.ReturnValues, "*"+entityName, "error")

	function.Body = "var " + variableName + " " + entityName + "\n\n" + "\tif err := r.db.Where(\"id = ?\", id).First(&" + variableName + ").Error; err != nil {" + "\n" + "\t\t" + "return nil, err" + "\n" + "\t}" + "\n\n" + "\treturn &" + variableName + ", nil"
	return function
}

// findByAllFields func where field = ?
func (g *Gorm) findByAllFields(info parse.StructInfo) []filesys_core.FuncBody {
	var functions []filesys_core.FuncBody

	for _, f := range info.Fields {
		var function filesys_core.FuncBody

		checkField := strings.ToLower(f.Name)

		if checkField == "id" || checkField == "createdat" || checkField == "updatedat" || checkField == "deletedat" {
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

		if unique {
			//	find one
			function = g.findOne(f, info.PackageName, info.Name)
			functions = append(functions, function)
		} else {
			//find many
			function = g.findMany(f, info.PackageName, info.Name)
			functions = append(functions, function)
		}

	}

	return functions
}

// findBY func where field = ?
func (g *Gorm) findOne(f parse.FieldInfo, packageName, name string) filesys_core.FuncBody {

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
func (g *Gorm) findMany(f parse.FieldInfo, packageName, name string) filesys_core.FuncBody {

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
