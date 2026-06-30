# Project Guidelines

## Overview

`gosimplectorgen` is a tiny Go code generator. Given a single `.go` source file containing
struct type declarations, it emits a companion `*_gen.go` file with a `NewXxx` constructor
per struct. Each constructor accepts every named field as a positional argument and returns
a pointer to a fully initialized struct literal. Typically wired via:

```go
//go:generate go tool gosimplectorgen $GOFILE
```

For installation, CLI flags, and a worked example, see [`README.md`](./README.md).

## Language

- Write all commit messages in English.
- Write all code comments in English.
- Write all pull request titles and descriptions in English.

## Snapshot Tests

Snapshot tests use [go-snaps](https://github.com/gkampitakis/go-snaps). To update snapshots, run:

```
UPDATE_SNAPS=true go test ./...
```
