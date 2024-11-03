package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/srinivassivaratri/RSSAggregator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	// Start immediate fetch, then use ticker for subsequent fetches
	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for {
		err := scrapeFeeds(s)
		if err != nil {
			log.Printf("error scraping feeds: %v", err)
		}
		<-ticker.C
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed: %w", err)
	}

	// Mark as fetched first to prevent excessive fetches on error
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	fmt.Printf("Fetching feed %s...\n", feed.Url)
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("Found %d posts in %s\n", len(rssFeed.Channel.Item), feed.Name)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}
	fmt.Println("------------------------")
	return nil
}
