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
