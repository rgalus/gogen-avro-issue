// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewBarWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewBar()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type BarReader struct {
	r io.Reader
	p *vm.Program
}

func NewBarReader(r io.Reader) (*BarReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewBar()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &BarReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r BarReader) Read() (*Bar, error) {
	t := NewBar()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
