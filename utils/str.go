package utils

import (
	"bufio"
	"strings"
)

func GetLines(content string) []string {
	reader := strings.NewReader(content)
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func FilterEmptyLines(lines []string) []string {
	var filteredLines []string
	for _, line := range lines {
		if strings.Trim(line, " \n\t") != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines
}
