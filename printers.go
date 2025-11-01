package main

import (
	"fmt"

	"github.com/alternateved/gator/internal/database"
)

func printUser(user database.User) {
	fmt.Printf("ID:      %s\n", user.ID)
	fmt.Printf("Name:    %v\n", user.Name)
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("ID:      %s\n", feed.ID)
	fmt.Printf("Created: %v\n", feed.CreatedAt)
	fmt.Printf("Updated: %v\n", feed.UpdatedAt)
	fmt.Printf("Name:    %s\n", feed.Name)
	fmt.Printf("URL:     %s\n", feed.Url)
	fmt.Printf("User:    %s\n", user.Name)

}
func printFeedFollow(userName, feedName string) {
	fmt.Printf("User:    %s\n", userName)
	fmt.Printf("Feed:    %s\n", feedName)
}
