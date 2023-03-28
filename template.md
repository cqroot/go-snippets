# Go Snippets

[![Action Status](https://github.com/cqroot/go-snippets/workflows/test/badge.svg)](https://github.com/cqroot/go-snippets/actions)

## Code Snippets

| #   | Type | Title |
| --- | ---- | ----- |

{{- range $index, $snippet := .Snippets }}
| {{ add $index 1 }} | {{ Title $snippet.Typ }} | [{{ Title $snippet.Title }}](https://github.com/cqroot/go-snippets/tree/main/snippets/{{ $snippet.Typ }}-{{ $snippet.Title }}/main.go) |
{{- end }}

## Design Patterns

| #   | Type | Title |
| --- | ---- | ----- |

{{- range $index, $pattern := .Patterns }}
| {{ add $index 1 }} | {{ Title $pattern.Typ }} Patterns | [{{ Title $pattern.Title }}](https://github.com/cqroot/go-snippets/tree/main/design-patterns/{{ $pattern.Typ }}-{{ $pattern.Title }}) |
{{- end }}
