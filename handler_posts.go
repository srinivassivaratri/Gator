package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/srinivassivaratri/RSSAggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("limit must be a number: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	fmt.Printf("Found %d posts:\n", len(posts))
	for _, post := range posts {
		fmt.Printf("\nTitle: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("Published: %v\n", post.PublishedAt.Time.Format("2006-01-02 15:04:05"))
		}
		fmt.Println("------------------------")
	}
	return nil
}
