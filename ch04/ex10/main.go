package main

import (
	"./github"
	"fmt"
	"log"
	"os"
	"time"
)

//!+
func main() {

	now := time.Now()
	beforeMonth := now.AddDate(0, -1, 0)
	beforeYear := now.AddDate(-1, 0, 0)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	//1月未満
	fmt.Printf("----------------------:\n")
	fmt.Printf("created at 1mo ~ now:\n")
	fmt.Printf("----------------------:\n")
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeMonth) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	//一ヶ月以上
	fmt.Printf("----------------------:\n")
	fmt.Printf("created between 1year ~ 1mo:\n")
	fmt.Printf("----------------------:\n")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeMonth) || 
				item.CreatedAt.Equal(beforeMonth) ||
				item.CreatedAt.After(beforeYear) {
					fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}

	//1年以上前
	fmt.Printf("----------------------:\n")
	fmt.Printf("created between ~ 1year:\n")
	fmt.Printf("----------------------:\n")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeYear) || 
				item.CreatedAt.Equal(beforeYear) {
					fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}