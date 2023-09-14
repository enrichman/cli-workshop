package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stargazer",
		Short: "Stargazer helps you starring Go repositories",
		Long: `
A very simple cli made during a workshop
that helps you searching and starring Go repositories.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			githubService := &GithubService{}

			user, err := githubService.GetUser(args[0])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("User '%v' (%v) found\n", user.Username, user.Name)
		},
	}
}
