# Go Snippets

[![Action Status](https://github.com/cqroot/go-snippets/workflows/test/badge.svg)](https://github.com/cqroot/go-snippets/actions)

{{ range $index, $snippet := . }}
{{ add $index 1 }}. [{{ name $snippet }}](https://github.com/cqroot/go-snippets/tree/main/{{ $snippet }}/main.go)
{{- end }}
