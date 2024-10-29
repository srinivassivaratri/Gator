// Declares the main package - entry point for the program
package main

// Import required packages for I/O and configuration handling
import (
	"fmt"
	"log"

	"github.com/srinivassivaratri/RSSAggregator/internal/config"
)

// Main function - program starts here
func main() {
	// Try to read the config file into cfg variable
	cfg, err := config.Read()
	// If reading fails, log the error and exit program
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	// Print the initial config values using %+v to show field names
	fmt.Printf("Initial read config: %+v\n", cfg)

	// Update username to "srinivas" and save to config file
	if err := cfg.SetUser("srinivas"); err != nil {
		// If saving fails, log error and exit
		log.Fatal(err)
	}

	// Read the config file again to verify changes
	updatedCfg, err := config.Read()
	// If reading fails, log error and exit
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	// Print the updated config to show changes
	fmt.Printf("Updated config: %+v\n", updatedCfg)
}
