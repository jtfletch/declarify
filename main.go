package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Function to parse toml file. All test implementation
func parseToml() interface{} {
	vi := viper.New()
	vi.SetConfigFile("test.yaml")
	vi.ReadInConfig()
	return vi.Get("arch.terminals")
}

func main() {

	// Use the Get() method to retrieve the value of the "terminals" key
	var terminals = parseToml()

	// Check if the value is a slice
	if terminalsSlice, ok := terminals.([]interface{}); ok {
		// Iterate over the slice and print each terminal
		for _, terminal := range terminalsSlice {
			fmt.Println(terminal)
		}
	} else {
		fmt.Println("Terminals key not found or not a slice")
	}

	// testing reading in from command line
	if len(os.Args) > 1 {
		// Print the value of the first command line argument
		fmt.Println("word:", os.Args[1])
	} else {
		fmt.Println("No command line argument provided")
	}
}
