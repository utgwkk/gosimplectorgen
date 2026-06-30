package testpkg

import "example.com/foo"

type myStruct struct {
	foo.Foo
	a int
	b string
}
