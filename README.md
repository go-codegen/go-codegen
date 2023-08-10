# go-codegen

[//]: # ([![GoDoc]&#40;https://godoc.org/github.com/kevinburke/go-codegen?status.svg&#41;]&#40;https://godoc.org/github.com/kevinburke/go-codegen&#41;)
## Overview
go-codegen is a library for generating Go code. It's useful for generating
boilerplate code, or code that is tedious to write by hand. It's similar to
[go-generate](https://blog.golang.org/generate), but go-codegen is a library
that you can use to generate code, instead of a tool that you run on the
command line.

## Installation

```
go get github.com/go-codegen/go-codegen

go install github.com/go-codegen/go-codegen
```

## Usage

[//]: # (TODO: add more examples)

### Create a entity 
````
type Recipient struct {
	gorm.Model
}

type User struct {
	gorm.Model
	NameAction string `gorm:"unique"`
	Age        int    `codegen:"unique"`
	RecipientID int
	Recipient
}
````



### Create a repository with gorm
````
go-codegen createRepository gorm --path=test/repository-test.go --out=test/files/
````

 

[//]: # (List of available commands)


 
## Tags

* unique - when you create a repository methods, the field will be unique in method find




