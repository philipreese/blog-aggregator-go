package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

    fmt.Println("Database reset successfully!")
	return nil
}