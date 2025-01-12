package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	characterLimit = 50000
	chaptersDir    = "chapters"
)

func main() {
	// Read all files in the chapters directory
	files, err := ioutil.ReadDir(chaptersDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".txt") {
			continue
		}

		// Read the content of each file
		filePath := filepath.Join(chaptersDir, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
			continue
		}

		// Remove square brackets from content
		cleanedContent := removeSquareBrackets(string(content))

		// Create output directory
		baseFileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		outputDir := filepath.Join(chaptersDir, baseFileName+"_parts")
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory for %s: %v\n", file.Name(), err)
			continue
		}

		// Split content into parts
		parts := splitContent(cleanedContent, characterLimit)

		// Write parts to separate files
		for i, part := range parts {
			outputPath := filepath.Join(outputDir, fmt.Sprintf("%s_part%d.txt", baseFileName, i+1))
			err = ioutil.WriteFile(outputPath, []byte(part), 0644)
			if err != nil {
				fmt.Printf("Error writing part %d of %s: %v\n", i+1, file.Name(), err)
				continue
			}
			fmt.Printf("Created %s\n", outputPath)
		}
	}
}

// Function to remove square brackets from text
func removeSquareBrackets(content string) string {
	return strings.ReplaceAll(strings.ReplaceAll(content, "[", ""), "]", "")
}

func splitContent(content string, limit int) []string {
	var parts []string
	contentRunes := []rune(content)
	totalLen := len(contentRunes)

	for i := 0; i < totalLen; i += limit {
		end := i + limit
		if end > totalLen {
			end = totalLen
		}

		// Try to find the last period or newline before the limit
		if end < totalLen {
			for j := end; j > i; j-- {
				if contentRunes[j-1] == '.' || contentRunes[j-1] == '\n' {
					end = j
					break
				}
			}
		}

		part := string(contentRunes[i:end])
		parts = append(parts, strings.TrimSpace(part))
	}

	return parts
}
