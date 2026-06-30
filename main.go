package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/utgwkk/gosimplectorgen/internal/generator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: gosimplectorgen <input.go>")
		os.Exit(2)
	}
	srcFile := os.Args[1]
	if !strings.HasSuffix(srcFile, ".go") {
		fmt.Fprintf(os.Stderr, "input file must end with .go: %s\n", srcFile)
		os.Exit(2)
	}
	outFile := strings.TrimSuffix(srcFile, ".go") + "_gen.go"

	f, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	g := generator.New(f, srcFile)
	if err := g.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "generate: %v\n", err)
		os.Exit(1)
	}
}
