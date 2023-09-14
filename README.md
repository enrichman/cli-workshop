# Section 05

In this section we are going to create two subcommands, a `user` command and add a new `search` command.

The user command is a copy paste of the `RootCommand`:

```go
func NewUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "user",
		Short: "User will get information of a github user",
		Args:  cobra.ExactArgs(1),
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
```

but now we will use this command as a subcommand of the root:

```go
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "stargazer",
		Short: "Stargazer helps you starring Go repositories",
		Long: `
A very simple cli made during a workshop
that helps you searching and starring Go repositories.`,
	}

	rootCmd.AddCommand(NewUserCmd())

	return rootCmd
}
```

Running a `go build -o stargazer .` and then `./stargazer` will now show the usage. If we want to fetch info of a user we need to run a `./stargazer user enrichman`:

```
-> % ./stargazer user enrichman
User 'enrichman' (Enrico Candino) found
```

Let's implement the `search` command.

Create a new command in the `cli.go`:

```go
func NewSearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Search will look for interesting repositories in Github",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
```

and add it to the Root command:

```go
rootCmd.AddCommand(
	NewUserCmd(),
	NewSearchCmd(),
)
```

and now we need to implement the Search in our GithubService.


Having a look at the Github documentation, and a sample response (https://api.github.com/search/repositories?q=Q) we can see that it will return a generic response with a list of repositories.

Let's create a generic `Repo` struct with the information that we need and the proper JSON tags

```go
type Repo struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	URL      string `json:"html_url"`
	Language string `json:"language"`
	Stars    int    `json:"stargazers_count"`
}
```

and the Search func

```go

func (s *GithubService) Search() ([]*Repo, error) {
	resp, err := http.Get("https://api.github.com/search/repositories?q=language:go")
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
```

We can now call the `Search` from the command:

```go
githubService := &GithubService{}

repos, err := githubService.Search()
if err != nil {
	log.Fatal(err)
}

for _, repo := range repos {
	fmt.Printf(
		"%-10d | %-30s | %-15s | %7d | %s\n",
		repo.ID, repo.FullName, repo.Language, repo.Stars, repo.URL,
	)
}
```

Let's improve our search adding the language flag.

Change the signature of the Search

```go
func (s *GithubService) Search(language string) ([]*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s", language)
	// code
}
```

and add the flag `--language` to the `search` command:

```go
func NewSearchCmd() *cobra.Command {
	var language string

	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search will look for interesting repositories in Github",
		Run: func(cmd *cobra.Command, args []string) {
			githubService := &GithubService{}

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
```