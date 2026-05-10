package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {
	fileToStream := filePath + ".processing"

	cmd := exec.Command(
		"ffmpeg",
		"-i", filePath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", fileToStream,
	)

	var buf bytes.Buffer
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("ffmpeg error: %s, %w", buf.String(), err)
	}

	return fileToStream, nil
}
