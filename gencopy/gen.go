package main

import (
	"log"
	"os"
	"text/template"
)

type GoType struct {
	Name, NameCamel, ZeroValue string
}

type GoTypes struct {
	Types []GoType
}

func main() {
	Types := GoTypes{[]GoType{
		{"string", "String", "\"\""},
		{"bool", "Bool", "false"},
		{"int32", "Int32", "0"},
		{"int64", "Int64", "0"},
		{"float32", "Float32", "0"},
		{"float64", "Float64", "0"},
	}}
	t, err := template.ParseFiles("templates/comparator.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("pkg/compare/compare.go")
	if err != nil {
		log.Fatal("create file: ", err)
	}
	err = t.Execute(f, Types)
	if err != nil {
		log.Fatal("execute: ", err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal("execute: ", err)
	}
}
