package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	configFilePath := flag.String("configuration", "config.yaml", "Path to the YAML configuration file")

	// Define a command-line flag for the root (folder path)
	var root string
	flag.StringVar(&root, "root", ".", "Root folder path")
	flag.Parse()

	// Validate if the root folder exists
	if err := validateFolderExists(root); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check if the file exists
	if _, err := os.Stat(*configFilePath); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist:", *configFilePath)
		return
	}

	// Read configuration from the file
	config, err := readConfigFromFile(*configFilePath)
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		return
	}

	fmt.Println("Name:", config.Version)

	fmt.Println("Namespaces:")
	for _, namespace := range config.Namespaces {
		fmt.Println("- Name:", namespace.Name)
	}
}
