package main

import (
	"context"
	"fmt"
)

// Deletes all users from database - useful for testing/reset
func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Successfully deleted all users and database is reset successfully")
	return nil
}