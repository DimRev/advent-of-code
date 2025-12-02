package lib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	RENDERER_DIR = "../render"
)

func PopulateRenderer(dayPart string, value int) error {
	rendererPath, err := filepath.Abs(RENDERER_DIR)
	if err != nil {
		return fmt.Errorf("error resolving renderer path: %v", err)
	}

	cmd := exec.Command("go", "run", "main.go", "populate", "go", dayPart, fmt.Sprintf("%d", value))
	cmd.Dir = rendererPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running populate command: %v", err)
	}

	return nil
}
