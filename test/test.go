package test

import (
	xz "github.com/go-codegen/go-codegen/test/files"
	"gorm.io/gorm"
)

type Hello struct {
	ID int
}

type RepositoryTest struct {
	gorm.Model
	HelloID int
	Hello
	WantHello  Hello
	NameAction string `codegen:"id" gorm:"unique"`
	Age        int    `codegen:"id,unique" json:"age" xml:"age"`
	Lol        xz.WhatThatIsFuck
}
