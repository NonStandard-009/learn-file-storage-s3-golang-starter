package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)

	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		return "", err
	}

	type fileDimensions struct {
		Streams []struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"streams"`
	}

	var fileDim fileDimensions
	if err := json.Unmarshal(buf.Bytes(), &fileDim); err != nil {
		return "", err
	}

	aspectRatio := fileDim.Streams[0].Width / fileDim.Streams[0].Height

	switch aspectRatio {
	case 1:
		return "16:9", nil
	case 0:
		return "9:16", nil
	default:
		return "other", nil
	}
}
