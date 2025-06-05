package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/philipreese/blog-aggregator-go/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeArg := cmd.Args[0]
	timeBetweenReq, err := time.ParseDuration(timeArg)
	if err != nil {
		return fmt.Errorf("invalid duration: %s - use format 1h 30m 15s or 3500ms", timeArg)
	}

	fmt.Printf("Collecting feeds every %s\n", timeArg)
	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("no new feeds to fetch: %w", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed `%s` as fetched: %w", feed.Name, err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't collect feed %s: %v", feed.Name, err)
	}

	for _, item := range rssFeed.Channel.Item {
		desc := sql.NullString{
			String: item.Description,
			Valid: true,
		}

		publishedAt := sql.NullTime{}
		if time, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time: time,
				Valid: true,
			}
		}
		
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Url: item.Link,
			Description: desc,
			PublishedAt: publishedAt,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "posts_url_key") {
				continue
			}
			log.Printf("couldn't create post: %v", err)
		}
	}

	fmt.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
	return nil
}