package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("err")
		return
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		fmt.Println("err")
		return
	}

	feedContent, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Printf("	--- %v ---\n", feedContent.Channel.Title)
	fmt.Printf("	%s\n", feedContent.Channel.Description)
	for i, item := range feedContent.Channel.Item {
		fmt.Printf("Item %d: %v\n", i, item.Title)
	}
}
