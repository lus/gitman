package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Read optional command line flags
	nameFlag := flag.String("name", "", "The name of the preset")
	pathFlag := flag.String("path", "", "The path to the preset file")
	flag.Parse()
	name := *nameFlag
	path := *pathFlag

	// Read the preset directory
	presets := os.Getenv("GITMAN_PRESETS")
	envEnabled := true
	if presets == "" {
		log.Println("INFO: Environment variable 'GITMAN_PRESETS' not found, disabling auto discover feature.")
		envEnabled = false
	}

	// Validate the input
	if !envEnabled && name != "" {
		log.Println("INFO: Ignoring 'name' flag because it is not relevant unless the 'GITMAN_PRESETS' environment variable was set.")
	}
	if envEnabled && name == "" && path == "" {
		name = "default"
	}
	if !envEnabled && path == "" {
		log.Println("ERROR: The 'path' flag is required if the 'GITMAN_PRESETS' environment variable was not set.")
		return
	}

	// Try to read the file
	filePath := presets + "\\" + name + ".txt"
	if path != "" {
		filePath = path
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("ERROR: Can't read file: " + err.Error())
		return
	}

	// Evaluate every line and set the corresponding git config variable
	for index, line := range strings.Split(string(file), "\n") {
		// Split the string into key and value
		split := strings.Split(line, "||")
		if len(split) != 2 {
			log.Println("ERROR: Failed to read line " + strconv.Itoa(index) + ": line must not contain more than one separator!")
			return
		}

		// Set the corresponding git config variable
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])
		err := exec.Command("git", "config", key, value).Run()
		if err != nil {
			log.Println("ERROR: Failed to set config variable to line " + strconv.Itoa(index) + ": " + err.Error())
			return
		}
		log.Println("INFO: Set git variable '" + key + "' to '" + value + "'.")
	}
}
