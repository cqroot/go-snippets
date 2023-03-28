# Go Snippets

[![Action Status](https://github.com/cqroot/go-snippets/workflows/test/badge.svg)](https://github.com/cqroot/go-snippets/actions)

| #   | Type | Title |
| --- | ---- | ----- |

{{- range $index, $snippet := . }}
| {{ add $index 1 }} | {{ ToUpper $snippet.Typ }} | [{{ Title $snippet.Title }}](https://github.com/cqroot/go-snippets/tree/main/{{ $snippet.Title }}/main.go) |
{{- end }}
