package example

import (
	"github.com/utgwkk/gosimplectorgen/example/bar"
	"github.com/utgwkk/gosimplectorgen/example/baz"
	"github.com/utgwkk/gosimplectorgen/example/foo"
)

//go:generate go tool gosimplectorgen $GOFILE

type privateStruct struct {
	foo foo.Foo
	bar bar.Bar
	baz baz.Baz
}
