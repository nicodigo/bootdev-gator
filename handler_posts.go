package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nicodigo/blog-aggregator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("usage: browse [limit]")
	}
	limit := 2
	if len(cmd.arguments) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
	}

	getPostsParams := database.GetPostsForUserParams{
		UserID:    user.ID,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println()
		fmt.Println("Post:", post.Title)
		fmt.Println("  - Description:", post.Description)
		if post.PublishedAt.Valid {
			fmt.Println("  - Pub. Date:", post.PublishedAt.Time)
		}
		fmt.Println("  - URL:", post.Url)
		fmt.Println()
	}
	return nil
}
