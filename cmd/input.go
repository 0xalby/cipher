package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readInput reads input from stdin(if available), a file or both.
func readInput(file string) (string, error) {
	var input strings.Builder
	// Checking if there's input from stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Reading from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			content := scanner.Text()
			// Add a newline only if content doesn't end with one
			if len(content) > 0 && content[len(content)-1] != '\n' {
				input.WriteString(content + "\n")
			}
			input.WriteString(content)
		}
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("error reading from stdin %s", err)
		}
	}
	// Reading from the file passed in
	if file != "" {
		content, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("error reading file %s %v", file, err)
		}
		input.WriteString(string(content))
		// Add a newline only if content doesn't end with one
		if len(content) > 0 && content[len(content)-1] != '\n' {
			input.WriteString("\n")
		}
	}
	//Returning an helpful message if no input is passed in
	if input.Len() == 0 {
		return "", fmt.Errorf("no input provided, please use stdin or pass in a file")
	}
	return input.String(), nil
}
