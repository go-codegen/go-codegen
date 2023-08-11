package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"
)

type ParsedField struct {
	Name         string
	Type         string
	Tags         map[string][]string
	NestedStruct *ParsedStruct
}

type ParsedStruct struct {
	StructName   string
	StructModule string
	FileName     string
	Fields       []ParsedField
}

type StructImpl struct {
	fset *token.FileSet
	f    *ast.File
}

func NewStruct() *StructImpl {
	return &StructImpl{}
}

func (s *StructImpl) parseFile(filename string) (*ast.File, error) {
	s.fset = token.NewFileSet()
	var err error
	s.f, err = parser.ParseFile(s.fset, filename, nil, parser.ParseComments)
	return s.f, err
}

// parse dir
func (s *StructImpl) parseDir(path string) ([]*ast.File, error) {

	//parse last part of path
	parts := strings.Split(path, "/")
	lastPart := parts[len(parts)-1]

	var pathWithoutLastPart string
	for i := 0; i < len(parts)-1; i++ {
		pathWithoutLastPart += parts[i] + "/"
	}

	//find files in path
	fInDir, err := os.ReadDir(pathWithoutLastPart)
	if err != nil {
		return nil, err
	}
	var files []*ast.File
	for _, file := range fInDir {
		if strings.Contains(file.Name(), lastPart) {
			newPath := pathWithoutLastPart + file.Name()
			fset := token.NewFileSet()
			pkgs, err := parser.ParseDir(fset, newPath, nil, parser.ParseComments)
			if err != nil {
				return nil, err
			}
			for _, pkg := range pkgs {
				for _, f := range pkg.Files {
					files = append(files, f)
				}
			}
		}
	}
	return files, nil
}

func (s *StructImpl) parseTags(tag string) map[string][]string {
	tags := make(map[string][]string)
	tag = strings.Trim(tag, "`")
	tagRegExp := regexp.MustCompile(`([\w]+):"([^"]+)"`)
	for _, t := range tagRegExp.FindAllStringSubmatch(tag, -1) {
		tags[t[1]] = strings.Split(t[2], ",")
	}
	return tags
}

func (s *StructImpl) findStructByName(structs [][]ParsedStruct, name string) []ParsedStruct {
	var result []ParsedStruct
	for _, innerSlice := range structs {
		for _, s := range innerSlice {
			if s.StructName == name {
				result = append(result, s)
			}
		}
	}
	return result
}

func (s *StructImpl) findFileByNameImport(name string) ([][]ParsedStruct, error) {
	for _, imp := range s.f.Imports {
		imp.Path.Value = strings.Trim(imp.Path.Value, "\"")
		parts := strings.Split(imp.Path.Value, "/")
		lastPart := parts[len(parts)-1]
		if lastPart == name {
			gopath := os.Getenv("GOPATH") // also consider "GOROOT" depending on your configuration
			if gopath == "" {
				log.Fatalf("No GOPATH set")
			}
			path := gopath + "/pkg/mod/" + imp.Path.Value
			path = strings.Replace(path, "\\", "/", -1)
			files, err := s.parseDir(path)
			if err != nil {
				return nil, err
			}
			var parsedStructs [][]ParsedStruct
			for _, f := range files {
				parsedStructs = append(parsedStructs, s.parseStructs(f.Name.Name, f))
			}

			return parsedStructs, nil
		}
	}
	return nil, fmt.Errorf("no file found")
}

func (s *StructImpl) collectFields(fields []*ast.Field) []ParsedField {
	var parsedFields []ParsedField
	for _, field := range fields {
		switch fieldType := field.Type.(type) {
		case *ast.StructType:
			// If the field type is a struct, recursively collect its fields
			if len(field.Names) > 0 {
				fieldName := field.Names[0].Name

				// Initialize parsed struct
				pstruct := ParsedStruct{StructName: fieldName}

				// Parse and collect fields from the nested struct
				pstruct.Fields = s.collectFields(fieldType.Fields.List)

				// Construct ParsedField with nested struct information
				parsedField := ParsedField{
					Name:         fieldName,
					Type:         "struct",
					NestedStruct: &pstruct,
				}

				parsedFields = append(parsedFields, parsedField)
			}
		case *ast.SelectorExpr:
			if fmt.Sprintf("%v", fieldType.X) == "time" {
				parsedFields = append(parsedFields, ParsedField{
					Name: field.Names[0].Name,
					Type: fmt.Sprintf("%v.%s", fieldType.X, fieldType.Sel.Name),
					NestedStruct: &ParsedStruct{
						StructName:   fieldType.Sel.Name,
						StructModule: fmt.Sprintf("%v", fieldType.X),
					},
				})
				continue
			}
			structures, err := s.findFileByNameImport(fmt.Sprintf("%v", fieldType.X))
			if err != nil {
				return nil
			}
			structs := s.findStructByName(structures, fieldType.Sel.Name)
			if len(structs) > 0 {
				parsedFields = append(parsedFields, ParsedField{
					Name: fmt.Sprintf("%v.%s", fieldType.X, fieldType.Sel.Name),
					Type: "struct",
					NestedStruct: &ParsedStruct{
						StructName:   fieldType.Sel.Name,
						StructModule: fmt.Sprintf("%v", fieldType.X),
						FileName:     structs[0].FileName,
						Fields:       structs[0].Fields,
					},
				})
				continue
			}
			// add struct to nested structs
			parsedFields = append(parsedFields, ParsedField{
				Name: fmt.Sprintf("%v.%s", fieldType.X, fieldType.Sel.Name),
				Type: "struct",
				NestedStruct: &ParsedStruct{
					StructName:   fieldType.Sel.Name,
					StructModule: fmt.Sprintf("%v", fieldType.X),
				},
			})
		case *ast.StarExpr:
			// If the field type is a pointer to another type
			if len(field.Names) > 0 && fieldType.X != nil {
				parsedFields = append(parsedFields, ParsedField{
					Name: field.Names[0].Name,
					Type: fmt.Sprintf("*%s", fieldType.X),
				})
			}
		case *ast.Ident:
			// If the field type is a struct, recursively collect its fields

			//TODO ADD SUPPORT FOR EMBEDDED STRUCTS

			if fieldType.Obj != nil {
				decl, ok := fieldType.Obj.Decl.(*ast.TypeSpec)
				if !ok {
					log.Fatal("Not a *ast.TypeSpec")
				}

				astStruct, ok := decl.Type.(*ast.StructType)
				if !ok {
					log.Fatal("Not a *ast.StructType")
				}

				fields := s.collectFields(astStruct.Fields.List)
				var name string
				if len(field.Names) > 0 {
					name = field.Names[0].Name
				} else {
					name = fieldType.Name
				}

				parsedFields = append(parsedFields, ParsedField{
					Name:         name,
					Type:         "struct",
					NestedStruct: &ParsedStruct{StructName: fieldType.Name, Fields: fields},
				})
				continue
			}
			// Assume this is a regular field, collect it ONLY if it has a name
			if len(field.Names) > 0 {
				tags := map[string][]string{}
				if field.Tag != nil {
					tags = s.parseTags(field.Tag.Value)
				}
				parsedFields = append(parsedFields, ParsedField{
					Name: field.Names[0].Name,
					Type: fieldType.Name,
					Tags: tags,
				})
				continue
			}

		default:
			continue
		}

	}
	return parsedFields
}

func (s *StructImpl) processSpecs(decl *ast.GenDecl, filename string) []ParsedStruct {
	var structs []ParsedStruct
	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		fields := s.collectFields(structType.Fields.List)

		structs = append(structs, ParsedStruct{FileName: filename, StructName: typeSpec.Name.Name, Fields: fields})
	}
	return structs
}

func (s *StructImpl) parseStructs(filename string, file *ast.File) []ParsedStruct {
	var structs []ParsedStruct
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		structs = append(structs, s.processSpecs(genDecl, filename)...)
	}
	return structs
}

func main() {
	filename := "test/test.go"
	s := StructImpl{}
	file, err := s.parseFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	structs := s.parseStructs(filename, file)
	for _, s := range structs {
		printStruct(s)
	}
}

// print struct
func printStruct(s ParsedStruct) {
	fmt.Printf("struct %s {\n", s.StructName)
	for _, f := range s.Fields {
		println("\t", f.Name, f.Type)
		for k, v := range f.Tags {
			println("\t\t", k)
			for _, vv := range v {
				println("\t\t\t", vv)
			}
		}
	}
	//nested
	fmt.Println("}")
}
