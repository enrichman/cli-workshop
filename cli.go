package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func NewRootCmd() (*cobra.Command, error) {
	githubService, err := NewGithubService("https://api.github.com")
	if err != nil {
		return nil, err
	}

	rootCmd := &cobra.Command{
		Use:   "stargazer",
		Short: "Stargazer helps you starring Go repositories",
		Long: `
A very simple cli made during a workshop
that helps you searching and starring Go repositories.`,
	}

	rootCmd.AddCommand(
		NewUserCmd(githubService),
		NewSearchCmd(githubService),
	)

	return rootCmd, nil
}

func NewUserCmd(githubService *GithubService) *cobra.Command {
	return &cobra.Command{
		Use:   "user",
		Short: "User will get information of a github user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			user, err := githubService.GetUser(args[0])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("User '%v' (%v) found\n", user.Username, user.Name)
		},
	}
}

func NewSearchCmd(githubService *GithubService) *cobra.Command {
	var language string

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search will look for interesting repositories in Github",
		Run: func(cmd *cobra.Command, args []string) {
			repos, err := githubService.Search(language)
			if err != nil {
				log.Fatal(err)
			}

			for _, repo := range repos {
				fmt.Printf(
					"%-10d | %-30s | %-15s | %7d | %s\n",
					repo.ID, repo.FullName, repo.Language, repo.Stars, repo.URL,
				)
			}
		},
	}

	searchCmd.Flags().StringVar(&language, "language", "", "search for repositories written in this language")

	return searchCmd
}
