package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nicodigo/blog-aggregator/internal/database"
)

func handleAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", *feed)
	return nil
}

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Usage: addfeed <name> <url>")
	}
	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currUser.ID,
	}

	addedFeed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}

	fmt.Println("Feed added to the database:")
	printFeedInfo(addedFeed, currUser)

	return createFeedFollows(s, currUser, addedFeed)
}

func handleFeeds(s *state, cmd command) error {
	feedList, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println(len(feedList), "Feeds found:")
	for _, feed := range feedList {
		feedUser, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		printFeedInfo(feed, feedUser)
	}
	return nil
}

func printFeedInfo(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}
