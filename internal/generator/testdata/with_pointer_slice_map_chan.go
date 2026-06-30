package testpkg

import (
	"example.com/bar"
	"example.com/foo"
)

type myStruct struct {
	a *foo.Foo
	b []bar.Bar
	c map[foo.Foo]bar.Bar
	d chan foo.Foo
}
