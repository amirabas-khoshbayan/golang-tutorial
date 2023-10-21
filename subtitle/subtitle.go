package main

import (
	"os"
	"path/filepath"
	"strings"
)

const subsDir = "subtitle/subs"

func main() {
	files, _ := os.ReadDir(subsDir)
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".srt") || strings.HasSuffix(fileName, ".vtt") {
			Process(fileName)
			break
		}
	}
}

func Process(filename string) {
	path := filepath.Join(subsDir, filename)
	bytes, _ := os.ReadFile(path)
	content := string(bytes)
	lines := strings.Split(content, "\n")

}
