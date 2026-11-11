package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nicodigo/blog-aggregator/internal/database"
)

func handleFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	return createFeedFollows(s, user, feed)
}

func createFeedFollows(s *state, user database.User, feed database.Feed) error {
	newFeedFollows := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	newRowFeedFollows, err := s.db.CreateFeedFollow(context.Background(), newFeedFollows)
	if err != nil {
		return err
	}

	fmt.Println("You are now following the feed:")
	fmt.Printf("Feed name: %s\n", newRowFeedFollows.FeedName)
	fmt.Printf("User name: %s\n", newRowFeedFollows.UserName)
	return  nil
}


func handleFollowing(s *state, cmd command) error {
	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	listFeedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currUser.ID)
	if err != nil {
		return err
	}

	fmt.Printf("You are following %d feeds:\n", len(listFeedFollows))
	for i, feedFollow := range listFeedFollows {
		fmt.Printf("%d. Feed name: %s\n", i, feedFollow.FeedName)
	}
	return  nil
}
