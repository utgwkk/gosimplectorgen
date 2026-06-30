package generator_test

import (
	"bytes"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/utgwkk/gosimplectorgen/internal/generator"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{name: "no_imports", input: "testdata/no_imports.go"},
		{name: "with_imports", input: "testdata/with_imports.go"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			g := generator.New(&buf, tc.input)
			if err := g.Generate(); err != nil {
				t.Fatalf("Generate: %v", err)
			}
			snaps.MatchSnapshot(t, buf.String())
		})
	}
}
