package test

import (
	"gorm.io/gorm"
)

type Hello struct {
	ID   int
	Data string
}

type RepositoryTest struct {
	gorm.Model
	HelloID int `codegen:"index"`
	Hello
	NameAction string `codegen:"id" gorm:"unique"`
	Age        int    `codegen:"id,unique" json:"age" xml:"age"`
}
