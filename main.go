package main

import (
	"html/template"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	files, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	snippets := make([]string, 0)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		snippets = append(snippets, file.Name())
	}

	tmpl, err := template.New("template.md").
		Funcs(FuncMap()).
		ParseFiles("template.md")
	if err != nil {
		panic(err)
	}

	fOutput, err := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		panic(err)
	}
	defer fOutput.Close()

	err = tmpl.Execute(fOutput, snippets)
	if err != nil {
		panic(err)
	}
}

func FuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"add": func(i int, j int) int {
			return i + j
		},
		"name": func(snippet string) string {
			parts := strings.Split(snippet, "_")
			if len(parts) < 2 {
				return snippet
			}
			name := strings.Replace(parts[1], "-", " ", -1)
			return cases.Title(language.English, cases.Compact).String(name)
		},
	}

	return funcMap
}
