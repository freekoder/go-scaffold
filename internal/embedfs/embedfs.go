package embedfs

import (
	"embed"
	"fmt"
)

//go:embed template/*
var template embed.FS

func GetFileTemplate(filename string) ([]byte, error) {
	return template.ReadFile(fmt.Sprintf("template/%s", filename))
}
