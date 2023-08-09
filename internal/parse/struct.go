package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

type Struct struct {
	Filename string
	Structs  []StructInfo
}

type StructInfo struct {
	Name        string
	Fields      []FieldInfo
	PackageName string
}

type FieldInfo struct {
	Name string
	Type string
	Tags map[string][]string
	Tag  string
}

func NewParse(filename string, prefix string) (*Struct, error) {
	r := &Struct{
		Filename: filename,
	}

	if err := r.parseStructs(prefix); err != nil {
		return nil, err
	}

	return r, nil
}

func (p *Struct) parseStructs(prefix string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, p.Filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					structName := typeSpec.Name.Name
					if strings.HasPrefix(structName, prefix) {
						fields, fieldErr := p.getFields(typeSpec.Type)
						if fieldErr != nil {
							return fieldErr
						}
						p.Structs = append(p.Structs, StructInfo{
							Name:        strings.TrimPrefix(structName, prefix),
							Fields:      fields,
							PackageName: node.Name.Name,
						})
					}
				}
			}
		}
	}
	return nil
}

func (p *Struct) getFields(node ast.Node) ([]FieldInfo, error) {
	var fields []FieldInfo
	if structType, ok := node.(*ast.StructType); ok {
		for _, field := range structType.Fields.List {

			if field.Names == nil {
				continue
			}

			fieldName := field.Names[0].Name
			fieldType := p.getTypeFromExpr(field.Type)

			var fieldTags map[string][]string
			if field.Tag != nil {
				var tagErr error
				fieldTags, tagErr = p.parseAllTagsInField(field.Tag.Value)

				if tagErr != nil {
					return nil, tagErr
				}
			}

			var tag string
			if field.Tag != nil {
				tag = field.Tag.Value
			}

			fields = append(fields, FieldInfo{
				Name: fieldName,
				Type: fieldType,
				Tags: fieldTags,
				Tag:  tag,
			})

		}
	}
	return fields, nil
}

func (p *Struct) parseAllTagsInField(tags string) (map[string][]string, error) {
	re := regexp.MustCompile(`(\w+):"([^"]+)"`)
	matches := re.FindAllStringSubmatch(tags, -1)

	keyValueMap := make(map[string][]string)
	for _, match := range matches {
		key := match[1]
		value := match[2]
		values := strings.Split(value, " ")
		keyValueMap[key] = values
	}

	return keyValueMap, nil

}

func (p *Struct) getTypeFromExpr(expr ast.Expr) string {
	switch exprType := expr.(type) {
	case *ast.Ident:
		return exprType.Name
	case *ast.StarExpr:
		return "*" + p.getTypeFromExpr(exprType.X)
	case *ast.SelectorExpr:
		return p.getTypeFromExpr(exprType.X) + "." + exprType.Sel.Name
	case *ast.ArrayType:
		return "[]" + p.getTypeFromExpr(exprType.Elt)
	case *ast.MapType:
		return "map[" + p.getTypeFromExpr(exprType.Key) + "]" + p.getTypeFromExpr(exprType.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.FuncType:
		return "func(...)"
	case *ast.ChanType:
		return "chan " + p.getTypeFromExpr(exprType.Value)
	case *ast.StructType:
		return "struct{}"
	default:
		return ""
	}
}
