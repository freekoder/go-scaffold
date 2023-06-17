package scaffold

import (
	"github.com/freekoder/go-scaffold/internal/config"
	"github.com/freekoder/go-scaffold/internal/embedfs"
	"html/template"
	"os"
)

func Build(outDir string, projectCfg config.Config) error {
	err := os.MkdirAll(outDir+"/"+projectCfg.Name, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(outDir+"/"+projectCfg.Name+"/cmd", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(outDir+"/"+projectCfg.Name+"/cmd/"+projectCfg.Name, os.ModePerm)
	if err != nil {
		return err
	}

	content, err := embedfs.GetFileTemplate("go.mod.tmpl")
	if err != nil {
		return err
	}

	goModTemplate, err := template.New("go.mod").Parse(string(content))
	if err != nil {
		return err
	}

	f, err := os.Create(outDir + "/" + projectCfg.Name + "/go.mod")

	err = goModTemplate.Execute(f, struct {
		ModuleName string
	}{ModuleName: projectCfg.Name})
	if err != nil {
		return err
	}
	f.Close()

	content, err = embedfs.GetFileTemplate("cmd/main.go")
	if err != nil {
		return err
	}

	err = os.WriteFile(outDir+"/"+projectCfg.Name+"/cmd/"+projectCfg.Name+"/"+projectCfg.Name+".go", content, 0644)
	if err != nil {
		return err
	}

	err = os.MkdirAll(outDir+"/"+projectCfg.Name+"/internal", os.ModePerm)
	if err != nil {
		return err
	}

	content, err = embedfs.GetFileTemplate("Makefile")
	if err != nil {
		return err
	}
	err = os.WriteFile(outDir+"/"+projectCfg.Name+"/Makefile", content, 0644)
	if err != nil {
		return err
	}
	return nil
}
