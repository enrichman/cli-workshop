package main

import (
	"fmt"
	"log"
)

func main() {
	githubService := &GithubService{}

	user, err := githubService.GetUser("enrichman")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User '%v' (%v) found\n", user.Username, user.Name)
}
