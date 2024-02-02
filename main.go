package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
	"os/exec"
  "path/filepath"
	"sort"
  "strings"
)

func ChangeWallpaper(file string) {
	// Separate the command and its arguments
	command := "nitrogen"
	args := []string{"--set-scaled", "/home/popo/Pictures/Wallpapers/" + file, "--head=0"}

	// Create the exec.Cmd object
	cmd := exec.Command(command, args...)
  fmt.Println(cmd)

	// Set the command's stdout to the current process's stdout
	cmd.Stdout = os.Stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

func main() {
	// Get the current directory
	currentDir, err := filepath.Abs("/home/popo/Pictures/Wallpapers") 
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// List and filter files in the current directory
	files, err := listJPGFiles(currentDir, "")
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}

	// Sort the files for a cleaner display
	sort.Strings(files)

	// Create a prompt with the file options
	prompt := promptui.Select{
		Label: "Select a file",
		Items: files,
	}

	// Prompt the user to select a file
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	// fmt.Println("You selected:", result)
	ChangeWallpaper(result)
}

func listJPGFiles(dir string, prefix string) ([]string, error) {
	var files []string

	// Read the directory
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Extract file names from fileInfos
	for _, fileInfo := range fileInfos {
		// Check if it's a directory
		if fileInfo.IsDir() {
			// Recursively list files in the subdirectory with the updated prefix
			subdir := filepath.Join(dir, fileInfo.Name())
			subfiles, err := listJPGFiles(subdir, filepath.Join(prefix, fileInfo.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, subfiles...)
		} else {
			// Check if the file has a ".jpg" extension
			if strings.HasSuffix(fileInfo.Name(), ".jpg") {
				// Add the file with the updated prefix to the list
				files = append(files, filepath.Join(prefix, fileInfo.Name()))
			}
		}
	}

	return files, nil
}

