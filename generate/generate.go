// Package main performs actual generation of the interval package functions for all built-in numeric types in Go.
// It is automatically called if `go generate` is issued from the project root.
package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	typeTemplate, err := template.ParseFiles("./generate/interval.tmpl")
	if err != nil {
		log.Fatalf("Template execution failed: %v", err)
	}
	testTemplate, err := template.ParseFiles("./generate/interval_test.tmpl")
	if err != nil {
		log.Fatalf("Template execution failed: %v", err)
	}
	types := []string{
		"int",
		"int64",
		"int32",
		"int16",
		"int8",
		"uint",
		"uint64",
		"uint32",
		"uint16",
		"uint8",
		"float32",
		"float64",
	}
	for _, t := range types {
		generateForType(typeTemplate, t, "./t_"+t+".go")
		generateForType(testTemplate, t, "./t_"+t+"_test.go")
	}
}

func generateForType(template *template.Template, typeName string, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("File creation failed: %v", err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()

	format := "Int"
	convert := "int64"
	bitSize := "0"
	if strings.HasPrefix(typeName, "uint") {
		format = "Uint"
		convert = "uint64"
	} else if strings.HasPrefix(typeName, "float") {
		format = "Float"
		convert = "float64"
		bitSize = typeName[5:] // 32 or 64
	}

	template.Execute(writer, map[string]interface{}{
		"T":           typeName,
		"CT":          strings.ToUpper(typeName[:1]) + typeName[1:],
		"FormatType":  format,
		"ConvertType": convert,
		"Unsigned":    format == "Uint",
		"Float":       format == "Float",
		"BitSize":     bitSize,
	})
}
