package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"text/template"
)

var reservedKeywords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true,
	"continue": true, "default": true, "defer": true, "else": true,
	"fallthrough": true, "for": true, "func": true, "go": true,
	"goto": true, "if": true, "import": true, "interface": true,
	"map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true,
	"var": true,
}

type DefaultField struct {
	Name  string
	Value string
}

type Field struct {
	Name      string
	Type      string
	JsonTag   string
	ParamName string
}

type Struct struct {
	Name           string
	Fields         []Field
	DefaultFields  []DefaultField
	BindingVersion string
	IsBinding      bool
	ParentBinding  bool
	NoMarshalFuncs bool
}

type TemplateData struct {
	Package    string
	Structs    []Struct
	AllStructs map[string]Struct
}

//TODO: migrate away from ast.Package and use go/types
func findParentBindings(pkg *ast.Package) map[string]bool {
	parentBindings := make(map[string]bool)

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				isBinding := false
				if genDecl.Doc != nil {
					for _, comment := range genDecl.Doc.List {
						if strings.Contains(comment.Text, "+binding") {
							isBinding = true
							break
						}
					}
				}

				if isBinding {
					for _, field := range structType.Fields.List {
						if starType, ok := field.Type.(*ast.StarExpr); ok {
							if ident, ok := starType.X.(*ast.Ident); ok {
								parentBindings[ident.Name] = true
							}
						}
					}
				}
			}
		}
	}
	return parentBindings
}

func findDefaultValues(field *ast.Field) []DefaultField {
	var defaults []DefaultField
	if field.Tag != nil {
		tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
		if defaultVal := tag.Get("default"); defaultVal != "" {
			defaults = append(defaults, DefaultField{
				Name:  field.Names[0].Name,
				Value: defaultVal,
			})
		}
	}
	return defaults
}

func getSafeParamName(jsonTag string) string {
	if jsonTag == "-" {
		return "additionalProperties"
	}
	paramName := strings.ReplaceAll(jsonTag, ".", "")
	if reservedKeywords[paramName] {
		return paramName + "Value"
	}
	return paramName
}

func processType(fieldType ast.Expr) (string, string, bool) {
	switch t := fieldType.(type) {
	case *ast.StarExpr:
		baseType, typeStr, isStruct := processType(t.X)
		return baseType, "*" + typeStr, isStruct
	case *ast.Ident:
		if t.Obj == nil {
			return t.Name, t.Name, false
		}
		typeSpec, ok := t.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return t.Name, t.Name, false
		}
		_, isStruct := typeSpec.Type.(*ast.StructType)
		return t.Name, t.Name, isStruct
	case *ast.ArrayType:
		baseType, typeStr, _ := processType(t.Elt)
		return baseType, "[]" + typeStr, false
	case *ast.MapType:
		keyType := fmt.Sprint(t.Key)
		valueType := fmt.Sprint(t.Value)
		if _, ok := t.Value.(*ast.InterfaceType); ok {
			valueType = "interface{}"
		}
		return "", fmt.Sprintf("map[%s]%s", keyType, valueType), false
	case *ast.InterfaceType:
		return "", "interface{}", false
	default:
		return "", fmt.Sprint(fieldType), false
	}
}

func processStruct(typeSpec *ast.TypeSpec, doc *ast.CommentGroup, allStructs map[string]Struct, parentBindings map[string]bool, seen map[string]bool) []Struct {
	structName := typeSpec.Name.Name
	if seen[structName] {
		return nil
	}
	seen[structName] = true

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil
	}

	isBinding := false
	noMarshalGen := false
	if doc != nil {
		for _, comment := range doc.List {
			if strings.Contains(comment.Text, "+binding") {
				isBinding = true
			}
			if strings.Contains(comment.Text, "+binding:marshal:no-gen") {
				noMarshalGen = true
			}
		}
	}

	mainStruct := Struct{
		Name:           structName,
		IsBinding:      isBinding,
		ParentBinding:  !isBinding && parentBindings[structName],
		NoMarshalFuncs: noMarshalGen,
	}

	var structs []Struct
	var fields []Field
	var defaultFields []DefaultField

	for _, field := range structType.Fields.List {
		if field.Tag == nil {
			continue
		}

		tag := strings.Trim(field.Tag.Value, "`")
		if !strings.Contains(tag, "json:\"") {
			continue
		}

		jsonTag := strings.Split(strings.Split(tag, "json:\"")[1], "\"")[0]
		jsonTag = strings.Split(jsonTag, ",")[0]

		_, fieldType, _ := processType(field.Type)
		fields = append(fields, Field{
			Name:      field.Names[0].Name,
			Type:      fieldType,
			JsonTag:   jsonTag,
			ParamName: getSafeParamName(jsonTag),
		})

		// Process default values
		defaults := findDefaultValues(field)
		defaultFields = append(defaultFields, defaults...)
	}

	mainStruct.Fields = fields
	mainStruct.DefaultFields = defaultFields
	if isBinding || mainStruct.ParentBinding {
		structs = append([]Struct{mainStruct}, structs...)
	}

	return structs
}

func main() {
	fset := token.NewFileSet()
	pkgDir := "."

	packages, err := parser.ParseDir(fset, pkgDir, func(fi os.FileInfo) bool {
		return !strings.HasPrefix(fi.Name(), "zz_generated")
	}, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	allStructs := make(map[string]Struct)
	parentBindings := make(map[string]bool)
	seen := make(map[string]bool)
	var finalStructs []Struct
	var pkgName string

	// First pass: collect all structs and parent bindings
	for _, pkg := range packages {
		pkgName = pkg.Name
		// Find parent bindings first
		for parent := range findParentBindings(pkg) {
			parentBindings[parent] = true
		}

		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if _, ok := typeSpec.Type.(*ast.StructType); ok {
								structs := processStruct(typeSpec, genDecl.Doc, allStructs, parentBindings, make(map[string]bool))
								if len(structs) > 0 {
									allStructs[typeSpec.Name.Name] = structs[0]
								}
							}
						}
					}
				}
			}
		}
	}

	// Second pass: process bindings and nested structs
	seen = make(map[string]bool)
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if _, ok := typeSpec.Type.(*ast.StructType); ok {
								structs := processStruct(typeSpec, genDecl.Doc, allStructs, parentBindings, seen)
								if len(structs) > 0 && (structs[0].IsBinding || structs[0].ParentBinding) {
									finalStructs = append(finalStructs, structs...)
								}
							}
						}
					}
				}
			}
		}
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}

	templatePath := filepath.Join(filepath.Dir(filename), "binding.go.tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	outputFile := filepath.Join(pkgDir, "zz_generated.binding.go")
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := TemplateData{
		Package:    pkgName,
		Structs:    finalStructs,
		AllStructs: allStructs,
	}

	if err := tmpl.Execute(f, data); err != nil {
		log.Fatal(err)
	}
}
