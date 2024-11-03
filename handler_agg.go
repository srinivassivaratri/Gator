package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"

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
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		// Try to create post, ignore if URL already exists
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				log.Printf("couldn't create post: %v", err)
			}
		}
	}
	fmt.Println("------------------------")
	return nil
}
