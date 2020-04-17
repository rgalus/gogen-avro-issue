package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/actgardner/gogen-avro/v7/compiler"

	"github.com/rgalus/gogen-avro-issue/avro"
)

func main() {
	pwd, _ := os.Getwd()
	schema, _ := ioutil.ReadFile(filepath.Join(pwd, "schema.avsc"))

	test := avro.DemoSchema{}

	_, err := compiler.CompileSchemaBytes(schema, []byte(test.Schema()))

	if err != nil {
		log.Fatal(err)
	}
}
