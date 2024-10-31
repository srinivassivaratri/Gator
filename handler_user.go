package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/srinivassivaratri/RSSAggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	// First check if user exists
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user %s already exists", name)
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("error checking user: %v", err)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting current user: %v", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user %s does not exist", name)
		}
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}



