package files

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func Read(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func IsJSON(filename string) error {
	if strings.ToLower(filepath.Ext(filename)) != ".json" {
		return errors.New("файл не является JSON")
	}
	return nil
}
