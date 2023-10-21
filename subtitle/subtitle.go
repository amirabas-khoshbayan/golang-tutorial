package main

import (
	"fmt"
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
	words := ExtractWords(lines)
	wordsMap := MakeWordsMap(words)
	fmt.Println(wordsMap)
}

func ExtractWords(lines []string) []string {
	words := []string{}
	indexer := 0
	for _, line := range lines {
		if line == "\r" {
			indexer += 2
			continue
		}
		if indexer > 0 {
			indexer--
			continue
		}

		newWords := strings.Split(line, " ")
		words = append(words, newWords...)
	}
	return words
}

func MakeWordsMap(words []string) map[string]int {
	wordsMap := make(map[string]int, 0)

	for _, word := range words {
		word := strings.NewReplacer(",", "", ".", "", "\r", "", "\"").Replace(word)

		if _, ok := wordsMap[word]; !ok {
			wordsMap[word] = 0
		}
		wordsMap[word]++
	}
	return wordsMap
}
