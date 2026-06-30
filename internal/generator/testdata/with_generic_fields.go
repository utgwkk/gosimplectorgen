package testpkg

import (
	"example.com/optional"
	"example.com/result"
)

type myStruct struct {
	a optional.Optional[string]
	b result.Result[int, error]
}
