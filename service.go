package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GithubService struct {
	URL string
}

func NewGithubService(baseURL string) (*GithubService, error) {
	if _, err := url.Parse(baseURL); err != nil {
		return nil, fmt.Errorf("creating GithubService: %w", err)
	}

	return &GithubService{
		URL: baseURL,
	}, nil
}

type User struct {
	Username string `json:"login"`
	Name     string `json:"name"`
}

func (s *GithubService) GetUser(username string) (*User, error) {
	endpoint := fmt.Sprintf("/users/%s", username)
	user := &User{}

	err := s.get(endpoint, user)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
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
	type SearchResponse struct {
		TotalCount int     `json:"total_count"`
		Items      []*Repo `json:"items"`
	}

	searchResponse := &SearchResponse{}

	q := url.Values{}
	q.Add("q", "language:"+language)

	url := fmt.Sprintf("/search/repositories?%s", q.Encode())
	err := s.get(url, searchResponse)
	if err != nil {
		return nil, fmt.Errorf("searching repositories: %w", err)
	}

	fmt.Printf("found %d repositories\n", searchResponse.TotalCount)

	return searchResponse.Items, nil
}

func (s *GithubService) get(endpoint string, response any) error {
	url := fmt.Sprintf("%s%s", s.URL, endpoint)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("executing get: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading the body response: %w", err)
	}

	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return fmt.Errorf("unmarshalling the body: %w", err)
	}
	return nil
}
