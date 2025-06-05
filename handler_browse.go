package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/philipreese/blog-aggregator-go/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		givenLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("usage: %s [limit]", cmd.Name)
		}
		limit = givenLimit
	}
	
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error retrieving posts for user: %w", err)
	}

	fmt.Printf("Found %v posts for user %s\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %s\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}