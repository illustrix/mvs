package auto

import (
	"fmt"
	"os"
	"path/filepath"
)

func getWorkingDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for dir != "/" {
		if _, err := os.Stat(dir + "/mvs.config.yaml"); err == nil {
			return dir, nil
		}
		dir = filepath.Dir(dir)
	}
	return "", fmt.Errorf("mvs.config.yaml not found in any parent directory")
}
