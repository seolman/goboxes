package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wrapInBox(text string) string {
	lines := strings.Split(text, "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	var box strings.Builder
	box.WriteString("+" + strings.Repeat("-", maxLen+2) + "+\n")
	for _, line := range lines {
		box.WriteString("| " + line + strings.Repeat(" ", maxLen-len(line)) + " |\n")
	}
	box.WriteString("+" + strings.Repeat("-", maxLen+2) + "+\n")

	return box.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder

	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
	inputText := strings.TrimSuffix(input.String(), "\n")
	output := wrapInBox(inputText)
	fmt.Print(output)
}
