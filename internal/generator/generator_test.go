package generator_test

import (
	"bytes"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/utgwkk/gosimplectorgen/internal/generator"
)

func TestGenerate(t *testing.T) {
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
		{name: "with_mixed_imports", input: "testdata/with_mixed_imports.go"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			g := generator.New(&buf, tc.input, tc.targetTypes)
			if err := g.Generate(); err != nil {
				t.Fatalf("Generate: %v", err)
			}
			snaps.MatchSnapshot(t, buf.String())
		})
	}
}
