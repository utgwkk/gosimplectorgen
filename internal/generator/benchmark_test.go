package generator_test

import (
	"io"
	"testing"

	"github.com/utgwkk/gosimplectorgen/internal/generator"
)

func BenchmarkGenerate(b *testing.B) {
	cases := []struct {
		name        string
		input       string
		targetTypes []string
	}{
		{name: "no_imports", input: "testdata/no_imports.go"},
		{name: "with_imports", input: "testdata/with_imports.go"},
		{name: "filter_types", input: "testdata/multiple_structs.go", targetTypes: []string{"foo", "baz"}},
		{name: "unused_imports_removed", input: "testdata/with_unused_imports.go"},
		{name: "with_generic_fields", input: "testdata/with_generic_fields.go"},
		{name: "with_func_fields", input: "testdata/with_func_fields.go"},
		{name: "with_pointer_slice_map_chan", input: "testdata/with_pointer_slice_map_chan.go"},
		{name: "with_embedded_fields", input: "testdata/with_embedded_fields.go"},
		{name: "exported_and_multiname", input: "testdata/exported_and_multiname.go"},
		{name: "with_aliased_imports", input: "testdata/with_aliased_imports.go"},
		{name: "non_struct_types", input: "testdata/non_struct_types.go"},
		{name: "multiple_structs_all", input: "testdata/multiple_structs.go"},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				g := generator.New(io.Discard, tc.input, tc.targetTypes)
				if err := g.Generate(); err != nil {
					b.Fatalf("Generate: %v", err)
				}
			}
		})
	}
}
