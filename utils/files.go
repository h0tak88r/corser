package utils

import (
	"os"
	
	"bufio"
	"strings"
	"github.com/zomasec/logz"

)

var (
	logger = logz.DefaultLogs()
)

func ReadFileLines(fileName string ) []string {
	lines := make([]string, 0)
	
	file, err := os.Open(fileName)
	
	if err != nil {
		logger.ERROR("Failed to open file: %s\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, strings.TrimSpace(scanner.Text()))
		}
		
	}

	if err := scanner.Err(); err != nil {
		logger.ERROR("Failed to read file: %s\n", err)
		
	}
	return lines
}
