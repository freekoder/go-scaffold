package scaffold

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/freekoder/go-scaffold/internal/config"
	"github.com/freekoder/go-scaffold/internal/embedfs"
	"text/template"
)

func Build(outDir string, projectCfg config.Config) error {
	basePath := createBasePath(outDir, projectCfg.Name)
	root := NewRoot(basePath)

	err := root.Init()
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}
	err = root.MkDir("internal")
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}
	mainContent, err := embedfs.ReadFile("cmd/main.go")
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}
	mainFile := fmt.Sprintf("cmd/service/%s.go", projectCfg.Name)
	err = root.WriteFile(mainFile, mainContent)
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}

	goModTemplate, err := embedfs.ReadTemplate("go.mod.tmpl")
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}

	goModContent, err := execTemplate(goModTemplate, struct {
		ModuleName string
	}{ModuleName: projectCfg.Name})
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}
	err = root.WriteFile("go.mod", goModContent)
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}

	makefileTemplate, err := embedfs.ReadTemplate("Makefile.tmpl")
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}
	makefileContent, err := execTemplate(makefileTemplate, struct {
		ModuleName string
	}{ModuleName: projectCfg.Name})
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}
	err = root.WriteFile("Makefile", makefileContent)
	if err != nil {
		return fmt.Errorf("init project: %v", err)
	}

	return nil
}

func execTemplate(tpl *template.Template, data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	bw := bufio.NewWriter(&buf)
	err := tpl.Execute(bw, data)
	if err != nil {
		return nil, err
	}
	_ = bw.Flush()
	return buf.Bytes(), nil
}

func createBasePath(outDir string, projectName string) string {
	return fmt.Sprintf("%s/%s", outDir, projectName)
}
