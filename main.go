package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/utgwkk/gosimplectorgen/internal/generator"
)

func main() {
	out := flag.String("out", "", "output file (default: <input>_gen.go)")
	targetTypes := flag.String("targetTypes", "", "comma-separated list of struct type names to generate constructors for (default: all structs)")
	version := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		return
	}

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: gosimplectorgen [-out <output.go>] [-targetTypes <Type1,Type2>] <input.go>")
		os.Exit(2)
	}
	srcFile := flag.Arg(0)
	if !strings.HasSuffix(srcFile, ".go") {
		fmt.Fprintf(os.Stderr, "input file must end with .go: %s\n", srcFile)
		os.Exit(2)
	}
	outFile := *out
	if outFile == "" {
		outFile = strings.TrimSuffix(srcFile, ".go") + "_gen.go"
	}

	var types []string
	if *targetTypes != "" {
		types = strings.Split(*targetTypes, ",")
	}

	f, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create output: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	g := generator.New(f, srcFile, types)
	if err := g.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "generate: %v\n", err)
		os.Exit(1)
	}
}
