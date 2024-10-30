// Declares the main package - entry point for the program
package main

// Import required packages for I/O and configuration handling
import (
	"fmt"
	"log"
	"os"

	"github.com/srinivassivaratri/RSSAggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

// Main function - program starts here
func main() {
	// Try to read the config file into cfg variable
	cfg, err := config.Read()
	// If reading fails, log the error and exit program
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	// Get command line args (excluding program name)
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1) // Exit with status code 1 when no command provided
	}

	// Create command from args
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Run the command
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}
