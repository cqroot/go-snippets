package main

import (
	"html/template"
	"os"
	"strings"

	. "github.com/cqroot/go-snippets"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Snippet struct {
	Typ   string
	Title string
}

type Pattern struct {
	Typ   string
	Title string
}

type Data struct {
	Snippets []Snippet
	Patterns []Pattern
}

func Snippets() []Snippet {
	files, err := os.ReadDir("./snippets")
	CheckErr(err)

	snippets := make([]Snippet, 0)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if !strings.Contains(file.Name(), "-") {
			continue
		}

		snippets = append(snippets, Snippet{
			Typ:   strings.Split(file.Name(), "-")[0],
			Title: strings.Split(file.Name(), "-")[1],
		})
	}

	return snippets
}

func Patterns() []Pattern {
	files, err := os.ReadDir("./design-patterns")
	CheckErr(err)

	patterns := make([]Pattern, 0)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if !strings.Contains(file.Name(), "-") {
			continue
		}

		patterns = append(patterns, Pattern{
			Typ:   strings.Split(file.Name(), "-")[0],
			Title: strings.Split(file.Name(), "-")[1],
		})
	}

	return patterns
}

func main() {
	tmpl, err := template.New("template.md").
		Funcs(FuncMap()).
		ParseFiles("template.md")
	CheckErr(err)

	fOutput, err := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	CheckErr(err)
	defer fOutput.Close()

	err = tmpl.Execute(fOutput, Data{
		Snippets: Snippets(),
		Patterns: Patterns(),
	})
	CheckErr(err)
}

func FuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"add": func(i int, j int) int {
			return i + j
		},
		"ToUpper": func(s string) string {
			return strings.ToUpper(s)
		},
		"Title": func(s string) string {
			name := strings.ReplaceAll(s, "_", " ")
			return cases.Title(language.English, cases.Compact).String(name)
		},
	}

	return funcMap
}
