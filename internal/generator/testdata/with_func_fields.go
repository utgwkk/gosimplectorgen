package testpkg

import (
	"example.com/bar"
	"example.com/foo"
)

type myStruct struct {
	handler func(foo.Foo) bar.Bar
	cb      func() error
}
