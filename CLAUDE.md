# Project Guidelines

## Language

- Write all commit messages in English.
- Write all code comments in English.
- Write all pull request titles and descriptions in English.

## Snapshot Tests

Snapshot tests use [go-snaps](https://github.com/gkampitakis/go-snaps). To update snapshots, run:

```
UPDATE_SNAPS=true go test ./...
```
