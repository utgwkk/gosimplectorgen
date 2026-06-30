package generator_test

import (
	"bytes"
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/utgwkk/gosimplectorgen/internal/generator"
)

var update = flag.Bool("update", false, "update golden files")

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

			goldenPath := strings.TrimSuffix(tc.input, ".go") + ".golden"
			if *update {
				if err := os.WriteFile(goldenPath, buf.Bytes(), 0644); err != nil {
					t.Fatalf("write golden: %v", err)
				}
			}

			want, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("read golden: %v", err)
			}

			if got := buf.Bytes(); !bytes.Equal(got, want) {
				t.Errorf("output mismatch\ngot:\n%s\nwant:\n%s", got, want)
			}
		})
	}
}
