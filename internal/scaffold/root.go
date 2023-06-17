package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
)

type Root struct {
	basePath string
}

func NewRoot(basePath string) *Root {
	return &Root{basePath: basePath}
}

func (r *Root) Init() error {
	return os.MkdirAll(r.basePath, os.ModePerm)
}

func (r *Root) MkDir(dirName string) error {
	newPath := fmt.Sprintf("%s/%s", r.basePath, dirName)
	return os.MkdirAll(newPath, os.ModePerm)
}

func (r *Root) WriteFile(filePath string, content []byte) error {
	fullPath := fmt.Sprintf("%s/%s", r.basePath, filePath)
	dirPath := filepath.Dir(fullPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(fullPath, content, os.ModePerm)
}
