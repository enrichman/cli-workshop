package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubService struct{}

type User struct {
	Username string `json:"login"`
	Name     string `json:"name"`
}

func (s *GithubService) GetUser(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reding the body response: %w", err)
	}

	user := &User{}
	err = json.Unmarshal(bodyBytes, user)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling the body: %w", err)
	}

	return user, nil
}

type Repo struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	URL      string `json:"html_url"`
	Language string `json:"language"`
	Stars    int    `json:"stargazers_count"`
}

func (s *GithubService) Search(language string) ([]*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s", language)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("searching repositories: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading the body response: %w", err)
	}

	type SearchResponse struct {
		TotalCount int     `json:"total_count"`
		Items      []*Repo `json:"items"`
	}

	searchResponse := &SearchResponse{}
	err = json.Unmarshal(bodyBytes, searchResponse)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling the body: %w", err)
	}

	fmt.Printf("found %d repositories\n", searchResponse.TotalCount)

	return searchResponse.Items, nil
}
