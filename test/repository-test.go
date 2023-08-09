package test

import "gorm.io/gorm"

type Hello struct {
	ID int
}
type RepositoryTest struct {
	gorm.Model
	HelloID int
	Hello
	NameAction string `codegen:"id" gorm:"unique"`
	Age        int    `codegen:"id,unique" json:"age" xml:"age"`
}
