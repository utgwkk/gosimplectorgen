package testpkg

import (
	"example.com/bar"
	"example.com/foo"
	"example.com/unused"
)

type myStruct struct {
	a foo.Foo
	b bar.Bar
}

var _ = new(unused.Something)
