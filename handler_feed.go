package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/philipreese/blog-aggregator-go/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user %s not found: %w", s.cfg.CurrentUserName, err)		
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	return nil
}

func printFeed(feed database.Feed, user database.User) {
    fmt.Printf("* ID:      %s\n", feed.ID);
    fmt.Printf("* Created: %s\n", feed.CreatedAt);
    fmt.Printf("* Updated: %s\n", feed.UpdatedAt);
    fmt.Printf("* name:    %s\n", feed.Name);
    fmt.Printf("* URL:     %s\n", feed.Url);
    fmt.Printf("* User:    %s\n", user.Name);
}