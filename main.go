package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var command, filename string
	var content []byte
	var isValid bool

	contentFromCommand := readFromStdin()

	// If there are 3 arguments, we assume command + filename
	if len(os.Args) == 3 {
		command = os.Args[1]
		filename = os.Args[2]
		content, isValid = checkIfFileExists(filename)
		if !isValid {
			return
		}
	} else if len(os.Args) == 2 && len(contentFromCommand) > 0 {
		// If there are 2 arguments and input from stdin is non-empty
		command = os.Args[1]
		content = contentFromCommand
	} else if len(os.Args) == 2 && len(contentFromCommand) == 0 {
		// Only filename provided, no piped input
		filename = os.Args[1]
		content, isValid = checkIfFileExists(filename)
		if !isValid {
			return
		}
		command = ""
	} else {
		// Neither a command nor valid input provided
		fmt.Println("Usage: ccwc [-c|-l|-w|-m] [filename] or piping input")
		return
	}

	// Process the command
	commandOutput(command, content, filename)
}

func readFromStdin() []byte {
	// Reading from stdin (for piped input like `cat file.txt | ccwc -l`)
	stat, _ := os.Stdin.Stat()

	// Only read from stdin if input is piped (not interactive)
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return nil
		}
		return stdinBytes
	}
	return nil
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
	case "": // Default: show lines, words, and bytes
		lines := numberOfLinesInAFile(file)
		words := numberOfWordsInAFile(file)
		bytesInAFile := numberOfBytesInAFile(file)
		fmt.Printf("%8d %8d %8d %s\n", lines, words, bytesInAFile, filename)
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
	return len([]rune(text)) // Multibyte character handling
}
