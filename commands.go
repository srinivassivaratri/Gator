package main

import "errors"

// command represents a CLI command with its arguments
type command struct {
	Name string   // Command name (e.g., "register")
	Args []string // Command arguments (e.g., ["username"])
}

// commands stores all available commands
type commands struct {
	// Maps command names to their handler functions
	registeredCommands map[string]func(*state, command) error
}

// register adds a new command handler
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// run executes a command if it exists
func (c *commands) run(s *state, cmd command) error {
	// Look up handler function
	handler, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return errors.New("command not found")
	}
	// Run handler with program state and command
	return handler(s, cmd)
}
