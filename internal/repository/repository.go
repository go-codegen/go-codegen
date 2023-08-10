package repository

import (
	"github.com/go-codegen/go-codegen/internal/colorPrint"
	"github.com/go-codegen/go-codegen/internal/filesys"
	filesys_core "github.com/go-codegen/go-codegen/internal/filesys/core"
	"github.com/go-codegen/go-codegen/internal/parse"
)

type Methods struct {
	Imports []string
	Struct  filesys_core.StructBody
	Funcs   []filesys_core.FuncBody
}

type Entity struct {
	Name    string
	Imports []string
	Struct  filesys_core.StructBody
}

type Module interface {
	MethodsData(info parse.StructInfo) Methods
}

type Repository struct {
	Module Module
	Struct *parse.Struct
	File   *filesys.FileSys
}

func NewRepository(module Module, structInfo *parse.Struct) *Repository {
	return &Repository{
		Module: module,
		Struct: structInfo,
		File:   filesys.NewFileSys(),
	}
}

func (r *Repository) Create(path string) error {
	if path == "" {
		path = "./"
	}
	for _, s := range r.Struct.Structs {
		////Entity
		//e := r.CreateEntity(s)
		//err := r.File.CreateFile(path+e.Name+".go", e)
		//if err != nil {
		//	colorPrint.PrintError( err)
		//}

		//Repository
		rm := r.CreateRepositoryMethods(s)
		err := r.File.CreateFile(path+rm.Name+".go", rm)
		if err != nil {
			colorPrint.PrintError(err)
			return err
		}
	}

	return nil
}

func (r *Repository) CreateRepositoryMethods(s parse.StructInfo) filesys_core.FileBody {

	var structBody filesys_core.FileBody

	repositoryData := r.Module.MethodsData(s)

	structBody.Name = r.File.CreateFileNameByStructName(s.Name, "repository-", "")

	structBody.Package = "repository"

	for _, i := range repositoryData.Imports {
		structBody.Imports = append(structBody.Imports, i)
	}

	structBody.Structs = append(structBody.Structs, repositoryData.Struct)

	for _, f := range repositoryData.Funcs {
		structBody.Funcs = append(structBody.Funcs, f)
	}

	return structBody
}

//func (r *Repository) CreateEntity(s parse.StructInfo) filesys_core.FileBody {
//
//	var structBody filesys_core.FileBody
//
//	entityData := r.Module.EntityData(s)
//
//	structBody.Name = r.File.CreateFileNameByStructName(s.Name, "entity-", "")
//
//	structBody.Package = s.Name + string(constants.Suffix)
//
//	for _, i := range entityData.Imports {
//		structBody.Imports = append(structBody.Imports, i)
//	}
//
//	structBody.Structs = append(structBody.Structs, entityData.Struct)
//
//	return structBody
//}
