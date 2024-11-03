package main

import (
	"context"
	"fmt"

	"github.com/srinivassivaratri/RSSAggregator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: agg")
	}

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Println("Feed Title:", feed.Channel.Title)
	fmt.Println("Feed Description:", feed.Channel.Description)
	fmt.Printf("Found %d posts:\n", len(feed.Channel.Item))
	for _, item := range feed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	return nil
}
