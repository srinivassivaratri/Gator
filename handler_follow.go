package main

import (
	"context"
	"fmt"
	"time"
	"database/sql"

	"github.com/google/uuid"
	"github.com/srinivassivaratri/RSSAggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <feed_url>")
	}

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("must be logged in to follow feeds. Use: login <username>")
	}

	url := cmd.Args[0]

	// Get current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	// Try to find feed by URL
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		if err == sql.ErrNoRows {
			// Feed doesn't exist, create it
			feed, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Name:      url, // Use URL as name initially
				Url:       url,
				UserID:    user.ID,
			})
			if err != nil {
				return fmt.Errorf("could not create feed: %w", err)
			}
		} else {
			return fmt.Errorf("could not check feed: %w", err)
		}
	}

	// Create feed follow
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("Following '%s' successfully!\n", feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: following")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get follows: %w", err)
	}

	fmt.Printf("You are following %d feeds:\n", len(follows))
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
} 