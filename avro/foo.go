// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type Foo struct {
	Id string `json:"id"`
}

const FooAvroCRC64Fingerprint = "h\xe3\x1c\x99\xd2\x0f\xf4\x00"

func NewFoo() *Foo {
	return &Foo{}
}

func DeserializeFoo(r io.Reader) (*Foo, error) {
	t := NewFoo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeFooFromSchema(r io.Reader, schema string) (*Foo, error) {
	t := NewFoo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeFoo(r *Foo, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Foo) Serialize(w io.Writer) error {
	return writeFoo(r, w)
}

func (r *Foo) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"foo\",\"type\":\"record\"}"
}

func (r *Foo) SchemaName() string {
	return "foo"
}

func (_ *Foo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Foo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Foo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Foo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Foo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Foo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Foo) SetString(v string)   { panic("Unsupported operation") }
func (_ *Foo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Foo) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	}
	panic("Unknown field index")
}

func (r *Foo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Foo) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Foo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Foo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Foo) Finalize()                        {}

func (_ *Foo) AvroCRC64Fingerprint() []byte {
	return []byte(FooAvroCRC64Fingerprint)
}
