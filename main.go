package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Config struct {
	CloneDir string `json:"cloneDir"`
	RepoURL  string `json:"repoURL"`
}

func main() {
	// Specify the path to the JSON configuration file
	configFile := "config.json"

	// Open and read the JSON configuration file
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %s", err)
	}
	defer file.Close()

	// Decode the JSON file into a Config struct
	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Failed to decode config file: %s", err)
	}

	// Run the 'git clone' command to clone the repository
	cmd := exec.Command("git", "clone", config.RepoURL, config.CloneDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to clone repository: %s", err)
	}

	fmt.Println("Source code retrieved successfully.")
}
