package embedfs

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed template/*
var embd embed.FS

func ReadFile(filename string) ([]byte, error) {
	return embd.ReadFile(fmt.Sprintf("template/%s", filename))
}

func ReadTemplate(filename string) (*template.Template, error) {
	content, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return template.New(filename).Parse(string(content))
}
