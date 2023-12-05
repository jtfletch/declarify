package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Function to parse toml file. All test implementation
func parseYAML() interface{} {
	vi := viper.New()
	vi.SetConfigFile("test.yaml")
	vi.SetConfigType("yaml")
	vi.ReadInConfig()
	return vi.Get("arch")
}

func main() {
	// Use the Get() method to retrieve the value of the "arch" key
	arch := parseYAML()

	// Check if the value is a map
	if archMap, ok := arch.(map[string]interface{}); ok {
		// Iterate over the map and print each subheading
		for key, value := range archMap {
			fmt.Printf("Subheading: %s\n", key)

			// Check if the subheading value is a slice
			if subheadingSlice, ok := value.([]interface{}); ok {
				// Iterate over the slice and print each element
				for _, element := range subheadingSlice {
					fmt.Println(element)
				}
			} else {
				fmt.Printf("Value under %s is not a slice\n", key)
			}
		}
	} else {
		fmt.Println("Arch key not found or not a map")
	}

	// Testing reading in from the command line
	if len(os.Args) > 1 {
		// Print the value of the first command line argument
		fmt.Println("Word:", os.Args[1])
	} else {
		fmt.Println("No command line argument provided")
	}
}
