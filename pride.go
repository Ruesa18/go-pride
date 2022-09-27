package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gookit/color"
)

type Configuration struct {
	FlagWidth  int
	FlagColors []int
}

func getDefaultConfig() Configuration {
	var conf Configuration
	conf.FlagWidth = 20
	conf.FlagColors = []int{160, 166, 220, 28, 4, 90}
	return conf
}

func displayFlag(flagWidth int, colors []uint8) {
	for i := 0; i < len(colors); i++ {
		for x := 0; x < flagWidth; x++ {
			color.S256(colors[i], colors[i]).Print(" ")
		}

		if i < len(colors)-1 {
			fmt.Println()
		}
	}
}

func main() {
	configuration := readConfig()

	arr := []uint8{}
	for i := 0; i < len(configuration.FlagColors); i++ {
		arr = append(arr, uint8(configuration.FlagColors[i]))
	}
	displayFlag(configuration.FlagWidth, arr)
}

func readConfig() Configuration {
	perm := fs.FileMode.Perm(0644)
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	directoryPath := filepath.Join(dirname, ".config", "go-pride")
	filePath := filepath.Join(directoryPath, "config.json")

	// Generate a default config.json if none exists
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		defaultConfig := getDefaultConfig()
		configJson, err := json.MarshalIndent(&defaultConfig, "", " ")

		if err != nil {
			log.Fatal("Error when creating JSON: ", err)
		}

		os.MkdirAll(directoryPath, perm)
		ioutil.WriteFile(filePath, configJson, perm)
	}

	// Read Configuration
	configFileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Configuration
	err = json.Unmarshal(configFileContent, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}
