package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func Resize(config ConfigResize, source, target, name string) (string, error) {
	sourcePath := filepath.Join(source, name)
	if !config.Enabled {
		return sourcePath, nil
	}

	targetName := strings.TrimSuffix(name, filepath.Ext(name)) + ".jpg"
	targetPath := filepath.Join(target, targetName)

	quality := fmt.Sprintf("%d", config.Quality)
	cmd := exec.Command("gm", "convert", sourcePath, "-quality", quality, targetPath)

	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err := cmd.Run()
	if err != nil {
		log.Printf("Resize: stdout: %s", cmd.Stdout)
		log.Printf("Resize: stderr: %s", cmd.Stderr)
		log.Printf("Resize: error: %v", err)
		return "", err
	}

	log.Printf("Resize: written as %s", targetPath)
	return targetPath, nil
}
