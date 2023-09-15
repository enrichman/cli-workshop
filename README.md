# Section 06

Organize the code. We see a lot of repetition in the `GithubService` so let's cleanup a bit.

We can start adding moving some common items to fields, and common operations in centralized functions.

Let's create a constructor that will accept a base URL, validating it. 

```go
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
```

Now we can create a `get` func that will handle the execution of the request, handling the response, and the unmarshaling of the body:

```go
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
```

this will simplify the `User` and the `Search`:

```go
func (s *GithubService) GetUser(username string) (*User, error) {
	endpoint := fmt.Sprintf("/users/%s", username)
	user := &User{}

	err := s.get(endpoint, user)
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}
	return user, nil
}
```

```go
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
```

Looking at the `cli.go` we see that we are creating a new GithubService in evry command. This can be improved handling the creation only in the RootCmd:

```go
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
```
