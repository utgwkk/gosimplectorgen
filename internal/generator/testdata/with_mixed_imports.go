package testpkg

import (
	"context"
	"time"

	"example.com/foo"
)

type myStruct struct {
	ctx     context.Context
	created time.Time
	item    foo.Foo
}
