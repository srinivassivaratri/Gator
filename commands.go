package main

import (
	"errors"
)

type command struct {
	// A command contains a name and a slice of string arguments
	Name string
	Args []string
}

type commands struct {
	// A map of command names to handler functions
	registeredCommands map[string]func(*state, command) error
}


func (c *commands) register(name string, f func(*state, command) error) {
	// This method registers a new handler function for a command name.
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	// This method runs a given command with the provided state if it exists
	handler, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return errors.New("command not found")
	}
	return handler(s, cmd)
}
