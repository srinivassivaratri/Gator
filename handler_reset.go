package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: reset")
	}

	// Reset config first
	s.cfg.CurrentUserName = ""
	err := s.cfg.SetUser("")
	if err != nil {
		return fmt.Errorf("error resetting config: %w", err)
	}

	// Then reset database
	err = s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database: %w", err)
	}

	fmt.Println("Database reset successfully")
	return nil
}
