package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Function to parse YAML file. All test implementation
func parseYAML() interface{} {
	vi := viper.New()
	vi.SetConfigFile("test.yaml")
	vi.SetConfigType("yaml")
	vi.ReadInConfig()
	return vi.Get("arch")
}

// Function to check if a value already exists in a slice
func contains(slice []interface{}, value interface{}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Function to append values to the tracker.yml file
func appendToYAML(values map[string][]interface{}) error {
	// Check if the tracker.yml file exists
	_, err := os.Stat("tracker.yml")
	newFile := os.IsNotExist(err)

	// Read existing data from tracker.yml if the file exists
	var existingData []byte
	if !newFile {
		existingData, err = os.ReadFile("tracker.yml")
		if err != nil {
			return err
		}
	}

	// Unmarshal existing data into a map with generic keys or create a new map if the file is new
	var data map[interface{}]interface{}
	if !newFile {
		err = yaml.Unmarshal(existingData, &data)
		if err != nil {
			return err
		}
	} else {
		data = make(map[interface{}]interface{})
	}

	// Ensure the "arch" key is present in the map
	if data["arch"] == nil {
		data["arch"] = make(map[interface{}]interface{})
	}

	// Append new entries to the existing data
	for key, value := range values {
		stringKey := fmt.Sprintf("%v", key) // Convert the key to a string
		if data["arch"].(map[interface{}]interface{})[stringKey] == nil {
			data["arch"].(map[interface{}]interface{})[stringKey] = value
		} else {
			// If the key already exists, append new entries to the existing slice
			existingSlice, ok := data["arch"].(map[interface{}]interface{})[stringKey].([]interface{})
			if !ok {
				return fmt.Errorf("existing value for key %s is not a slice", stringKey)
			}
			data["arch"].(map[interface{}]interface{})[stringKey] = append(existingSlice, value...)
		}
	}

	// Marshal the updated data
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	// Write the updated data back to tracker.yml
	err = os.WriteFile("tracker.yml", yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Use the Get() method to retrieve the value of the "arch" key
	arch := parseYAML()

	// Check if the value is a map
	if archMap, ok := arch.(map[string]interface{}); ok {
		// Check if a command line argument is provided
		if len(os.Args) > 1 {
			// Convert the command line argument to lowercase for case-insensitive comparison
			arg := strings.ToLower(os.Args[1])

			// Check if the command line argument matches any subheading
			if subheadingSlice, ok := archMap[arg].([]interface{}); ok {
				// Print elements under the matched subheading
				fmt.Printf("Subheading: %s\n", arg)
				for _, element := range subheadingSlice {
					fmt.Println(element)
				}

				// Append values to tracker.yml
				values := map[string][]interface{}{arg: subheadingSlice}
				err := appendToYAML(values)
				if err != nil {
					fmt.Println("Error appending to tracker.yml:", err)
				}
			} else if arg == "arch" {
				// Print all elements if the command line argument is "arch"
				for key, value := range archMap {
					fmt.Printf("Subheading: %s\n", key)
					if subheadingSlice, ok := value.([]interface{}); ok {
						for _, element := range subheadingSlice {
							fmt.Println(element)
						}

						// Append values to tracker.yml
						values := map[string][]interface{}{key: subheadingSlice}
						err := appendToYAML(values)
						if err != nil {
							fmt.Println("Error appending to tracker.yml:", err)
						}
					}
				}
			} else {
				fmt.Println("Please enter a valid subheading")
			}
		} else {
			fmt.Println("No command line argument provided")
		}
	} else {
		fmt.Println("Arch key not found or not a map")
	}
}
