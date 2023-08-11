package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

func parseFile(filename string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, filename, nil, parser.ParseComments)
}

func parseTags(tag string) map[string][]string {
	tags := make(map[string][]string)
	tag = strings.Trim(tag, "`")
	tagRegExp := regexp.MustCompile(`([\w]+):"([^"]+)"`)
	for _, t := range tagRegExp.FindAllStringSubmatch(tag, -1) {
		tags[t[1]] = strings.Split(t[2], ",")
	}
	return tags
}

func collectFields(fields []*ast.Field) []ParsedField {
	var parsedFields []ParsedField
	for _, field := range fields {
		fmt.Println(field.Type)
		switch fieldType := field.Type.(type) {
		case *ast.StructType:
			// If the field type is a struct, recursively collect its fields
			if len(field.Names) > 0 {
				fieldName := field.Names[0].Name

				// Initialize parsed struct
				pstruct := ParsedStruct{StructName: fieldName}

				// Parse and collect fields from the nested struct
				pstruct.Fields = collectFields(fieldType.Fields.List)

				// Construct ParsedField with nested struct information
				parsedField := ParsedField{
					Name:         fieldName,
					Type:         "struct",
					NestedStruct: &pstruct,
				}

				parsedFields = append(parsedFields, parsedField)
			}
		case *ast.SelectorExpr:
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
			if fieldType.Obj != nil {
				fields := collectFields(fieldType.Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List)

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
					tags = parseTags(field.Tag.Value)
				}
				parsedFields = append(parsedFields, ParsedField{
					Name: field.Names[0].Name,
					Type: fieldType.Name,
					Tags: tags,
				})
				continue
			}

		default:
			fmt.Println("Unknown field type", fieldType)
			//	print field type
			fmt.Println(fieldType)
		}

	}
	return parsedFields
}

func processSpecs(decl *ast.GenDecl, filename string) []ParsedStruct {
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
		fields := collectFields(structType.Fields.List)
		structs = append(structs, ParsedStruct{FileName: filename, StructName: typeSpec.Name.Name, Fields: fields})
	}
	return structs
}

func parseStructs(filename string, file *ast.File) []ParsedStruct {
	var structs []ParsedStruct
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		structs = append(structs, processSpecs(genDecl, filename)...)
	}
	return structs
}

func main() {
	filename := "test/test.go"
	file, err := parseFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("parse")
	structs := parseStructs(filename, file)
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
