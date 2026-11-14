package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nicodigo/blog-aggregator/internal/database"
)

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	feedContent, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("	--- %v ---\n", feedContent.Channel.Title)
	fmt.Printf("	%s\n", feedContent.Channel.Description)
	for _, item := range feedContent.Channel.Item {
		fmt.Printf("\nSaving %s...\n", item.Title)
		savePost(s, &item, nextFeed.ID)
	}
}

func savePost(s *state, feedItem *RSSItem, feedID uuid.UUID) {
	post := database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: feedItem.Title,
		Url: feedItem.Link,
		Description: feedItem.Description,
		PublishedAt: parsePublishedAt(feedItem.PubDate),
		FeedID: feedID,
	}
	_, err := s.db.CreatePost(context.Background(), post)
	if err != nil {
		fmt.Println("error saving post:", err)
	}
}

func parsePublishedAt(dateStr string) sql.NullTime {
    layouts := []string{
        time.RFC1123Z,
        time.RFC1123,
    }

    for _, layout := range layouts {
        t, err := time.Parse(layout, dateStr)
        if err == nil {
            return sql.NullTime{Time: t, Valid: true}
        }else {
			fmt.Println("error parsing the date:", err)
		}
    }

    return sql.NullTime{}
}
