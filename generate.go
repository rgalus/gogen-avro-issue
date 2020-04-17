package generate

//go:generate mkdir -p ./avro
//go:generate go run github.com/actgardner/gogen-avro/v7/cmd/gogen-avro -containers ./avro schema.avsc
