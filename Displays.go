package main

import (
	"fmt"
  "os/exec"
  "regexp"
)

// Search through string to return display port names...
func findConnectedWords(input string) []string {
	// Define a regular expression to find the word "connected" and capture the word before it
	re := regexp.MustCompile(`(\S+)\s+connected`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	// Extract the captured words from the matches
	var connectedWords []string
	for _, match := range matches {
		if len(match) >= 2 {
			connectedWords = append(connectedWords, match[1])
		}
	}

	return connectedWords
}

// Finds connected displays
func FindDisplays() []string {
  command := "xrandr" 
  args := []string{"--query"}
  displays := exec.Command(command, args...) 

  // Captures the output of the command
  output, err := displays.CombinedOutput()
  if err != nil {
    // Print error and return empty string
    fmt.Println("Error executing:", err)
    return []string{}
  }
  connectedDisplays := findConnectedWords(string(output))
 
  return connectedDisplays
}

