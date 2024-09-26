package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ccwc -[c|l|w|m] filename")
		return
	}

	command := os.Args[1]
	filename := os.Args[2]

	file, isValid := checkIfFileExists(filename)
	if !isValid {
		return
	}
	commandOutput(command, file, filename)
}

func commandOutput(command string, file []byte, filename string) {
	switch command {
	case "-c":
		fmt.Printf("%8d %s\n", numberOfBytesInAFile(file), filename)
	case "-l":
		fmt.Printf("%8d %s\n", numberOfLinesInAFile(file), filename)
	case "-w":
		fmt.Printf("%8d %s\n", numberOfWordsInAFile(file), filename)
	case "-m":
		fmt.Printf("%8d %s\n", numberOfCharactersInAFile(file), filename)
	default:
		fmt.Println("Invalid command. Use -c, -l, -w, or -m.")
	}
}

func checkIfFileExists(filename string) ([]byte, bool) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, false
	}
	return content, true
}

// Command -c: Number of bytes
func numberOfBytesInAFile(content []byte) int {
	return len(content)
}

// Command -l: Number of lines
func numberOfLinesInAFile(content []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	return lineCount
}

// Command -w: Number of words
func numberOfWordsInAFile(content []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	wordCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}
	return wordCount
}

// Command -m: Number of characters
func numberOfCharactersInAFile(content []byte) int {
	text := string(content)
	return len([]rune(text)) // Handling multibyte characters
}
