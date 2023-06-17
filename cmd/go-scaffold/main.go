package main

import (
	"fmt"
	"github.com/freekoder/go-scaffold/internal/config"
	"github.com/freekoder/go-scaffold/internal/scaffold"
	"os"
)

func main() {
	fmt.Print("Enter project name:")
	var projectName string
	_, err := fmt.Scanf("%s", &projectName)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(-1)
	}
	fmt.Printf("project name: %s\n", projectName)

	projectCfg := config.Config{
		Name: projectName,
	}

	outDir := "/tmp"

	err = scaffold.Build(outDir, projectCfg)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
