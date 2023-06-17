package scaffold

import (
	"github.com/freekoder/go-scaffold/internal/config"
	"github.com/freekoder/go-scaffold/internal/embedfs"
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
	err = os.MkdirAll(outDir+"/"+projectCfg.Name+"/internal", os.ModePerm)
	if err != nil {
		return err
	}

	content, err := embedfs.GetFileTemplate("Makefile")
	if err != nil {
		return err
	}
	err = os.WriteFile(outDir+"/"+projectCfg.Name+"/Makefile", content, 0644)
	if err != nil {
		return err
	}
	return nil
}
