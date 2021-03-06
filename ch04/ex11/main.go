package main

import (
	"fmt"
	"log"
	"os"

	"./github"
)

var usage string = `
search QUERY
[create|read|update|delete] OWNER REPO ISSUENUMBER
`

func exit() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args[:]) < 3 {
		exit()
	}
	
	fmt.Println(os.Args[1:])
	mode := os.Args[1]
	owner := os.Args[2]
	repo := os.Args[3]
	number := os.Args[4]
	
	//read functionality
	if mode == "read" {
		issue, err := github.GetIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
		body := issue.Body
		if body == "" {
			body = "<empty>\n"
		}
	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
		owner, repo, number, issue.User.Login, issue.Title, body)
	}
	
	//create functionality
	//deal with it later
	if mode == "create" {
		fmt.Println("Create")
		github.CreateIssue(owner, repo, number)
	}
	
	//search functionality
	if mode == "search" {
		result, err := github.SearchIssues(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n %s\n Time:%v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		
		}
	}
}
