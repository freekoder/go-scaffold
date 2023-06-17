package scaffold

import (
	"github.com/freekoder/go-scaffold/internal/config"
	"os"
)

func Build(outDir string, projectCfg config.Config) error {
	err := os.MkdirAll(outDir+"/"+projectCfg.Name, 0750)
	if err != nil {
		return err
	}
	return nil
}
