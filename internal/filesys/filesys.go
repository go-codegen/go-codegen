package filesys

import (
	"github.com/go-codegen/go-codegen/internal/filesys/core"
	"os"
	"path/filepath"
)

type FileSys struct {
	Body *filesys.Body
}

func NewFileSys() *FileSys {

	Func := filesys.NewFunc()
	Import := filesys.NewImport()
	Struct := filesys.NewStruct(Func)
	Interface := filesys.NewInterface()
	Body := filesys.NewBody(Func, Import, Struct, Interface)

	return &FileSys{
		Body: Body,
	}
}

func (f *FileSys) osCreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		if os.IsNotExist(err) {
			dir, _ := filepath.Split(filename)

			dirs := filepath.SplitList(dir)

			for _, d := range dirs {
				err := f.osCreateDir(d)
				if err != nil {
					return nil, err
				}
			}

			file, err = os.Create(filename)
			if err != nil {
				return nil, err
			}

			return file, nil
		}
		return nil, err
	}

	return file, nil
}

func (f *FileSys) osDeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileSys) osCreateDir(dir string) error {
	err := os.Mkdir(dir, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileSys) osDeleteDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileSys) CreateFile(path string, body filesys.FileBody) error {
	file, err := f.osCreateFile(path)
	if err != nil {
		return err
	}

	bodyFile := f.Body.Create(body)

	_, err = file.WriteString(bodyFile)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	//colorPrint.PrintSuccess(fmt.Sprintf("File %s created successfully.\n", path))
	return nil
}

func (f *FileSys) CreateDir(dir string) error {
	err := f.osCreateDir(dir)
	if err != nil {
		return err
	}
	//colorPrint.PrintSuccess(fmt.Sprintf("Dir %s created successfully.\n", dir))
	return nil
}

// CreateFileNameByStructName parse name of struct if letter is uppercase add - before letter and convert to lowercase
// example: UserTest -> user-test
// example: User -> user
func (f *FileSys) CreateFileNameByStructName(name string, prefix string, suffix string) string {
	var result string

	if prefix != "" {
		result += prefix
	}

	for i, s := range name {

		if i == 0 {
			result += string(s + 32)
			continue
		}

		if s >= 'A' && s <= 'Z' {
			//set to lowercase
			result += "-" + string(s+32)
			continue
		}

		result += string(s)
	}

	if suffix != "" {
		result += suffix
	}

	return result
}
